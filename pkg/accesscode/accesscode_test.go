package accesscode

import (
	"fmt"
	hobbyfarmiov1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	"github.com/hobbyfarm/gargantua/pkg/client/clientset/versioned/fake"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
	"testing"
	"time"
)

var accessCodeClient *AccessCodeClient

var accessCodes = []string{"test1", "test2", "test3"}
var expiredCodes = []string{"test1_expired", "test2_expired", "test3_expired"}
var badTimeCode = "bad_time_code"
var scenarios = []string{"scenario1", "scenario2"}
var courses = []string{"course1", "course2"}
var vmSets = []string{"vmset1", "vmset2"}
var restrictedBindValue = "rbvalue"

func TestMain(m *testing.M) {
	var err error
	fakeClient := fake.NewSimpleClientset() // fake clientset

	// create access codes
	for _, v := range accessCodes {
		ac := createAccessCode(v, v, scenarios, courses, time.Now().Add(time.Minute * 5).Format(time.UnixDate), vmSets, true, restrictedBindValue)
		_, err := fakeClient.HobbyfarmV1().AccessCodes().Create(&ac)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// create expired access codes
	for _, v := range expiredCodes {
		// time.Now() should be in the past once we query
		ac := createAccessCode(v, v, scenarios, courses, time.Now().Format(time.UnixDate), vmSets, true, restrictedBindValue)
		_, err := fakeClient.HobbyfarmV1().AccessCodes().Create(&ac)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// create access code with bad time value
	ac := createAccessCode(badTimeCode, badTimeCode, scenarios, courses, "lksjdflkjsdf", vmSets, true, restrictedBindValue)
	_, err = fakeClient.HobbyfarmV1().AccessCodes().Create(&ac)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// create user and associate access codes
	user := createUser("user", "user", "user@example.com", "password", accessCodes, true)
	_, err = fakeClient.HobbyfarmV1().Users().Create(&user)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	userExpiredAccessCodes := createUser("user_expired", "user_expired", "user_expired@example.com", "password", expiredCodes, true)
	_, err = fakeClient.HobbyfarmV1().Users().Create(&userExpiredAccessCodes)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	accessCodeClient, err = NewAccessCodeClient(fakeClient)
	if err != nil {
		fmt.Printf("error building access code client: %v \n", err)
		os.Exit(1)
	}
	
	os.Exit(m.Run())
}

func TestGetAccessCode_EmptyCode(t *testing.T) {
	_, err := accessCodeClient.GetAccessCode("", false)
	if err == nil {
		t.Error("did not error when 0 len code passed in")
	}
}

func TestGetAccessCode_BadCode(t *testing.T) {
	_, err := accessCodeClient.GetAccessCode("lkjsdflkjsdflkjsdf", false)
	if err == nil {
		t.Error("did not error when invalid access code passed in")
	}
}

func TestGetAccessCode_ExpiredNotOK(t *testing.T) {
	for _, v := range accessCodes {
		ac, err := accessCodeClient.GetAccessCode(v, false)
		if err != nil {
			t.Errorf("error getting access code %s: %v", v, err)
		}

		if ac.Spec.Code != v {
			t.Errorf("string value %s does not match access code %s", v, ac.Spec.Code)
		}
	}
}

func TestGetAccessCode_ExpiredOK(t *testing.T) {
	for _, v := range accessCodes {
		ac, err := accessCodeClient.GetAccessCode(v + "_expired", true)
		if err != nil {
			t.Errorf("error getting access code %s: %v", v, err)
		}

		if ac.Spec.Code != v {
			t.Errorf("string value %s does not match access code %s", v, ac.Spec.Code)
		}
	}
}

func TestGetAccessCodes_ExpiredOK(t *testing.T) {
	expiredCodes := make([]string, len(accessCodes))
	for i, v := range accessCodes {
		expiredCodes[i] = v + "_expired"
	}

	acList, err := accessCodeClient.GetAccessCodes(expiredCodes, true)
	if err != nil {
		t.Error(err)
	}

	for i, v := range acList {
		if v.Spec.Code != expiredCodes[i] {
			t.Errorf("string value %s does not match access code %s", accessCodes[i], v.Spec.Code)
		}
	}
}

func TestGetAccessCodes_ExpiredNotOK(t *testing.T) {
	acList, err := accessCodeClient.GetAccessCodes(accessCodes, false)
	if err != nil {
		t.Error(err)
	}

	for i, v := range acList {
		if v.Spec.Code != accessCodes[i] {
			t.Errorf("string value %s does not match access code %s", accessCodes[i], v.Spec.Code)
		}
	}
}

func TestGetAccessCodes_EmptyList(t *testing.T) {
	_, err := accessCodeClient.GetAccessCodes([]string{}, false)
	if err == nil {
		t.Errorf("did not error when passed empty access code list")
	}
}

func TestGetAccessCodes_BadTime(t *testing.T) {
	_, err := accessCodeClient.GetAccessCodes([]string{badTimeCode}, false)
	if err == nil {
		t.Errorf("did not error when access code contained invalid time")
	}
}

func TestGetScenarioIds(t *testing.T) {
	ids, err := accessCodeClient.GetScenarioIds(accessCodes[0])
	if err != nil {
		t.Error(err)
	}

	if len(ids) < 1 {
		t.Error("no scenario ids returned for valid access code")
	}

	var valid = true
	for i, v := range ids {
		if scenarios[i] != v {
			valid = false
		}
	}

	if !valid {
		t.Errorf("invalid scenarios returned, expected %s, got %s", scenarios, ids)
	}
}

func TestGetScenarioIds_EmptyCode(t *testing.T) {
	_, err := accessCodeClient.GetScenarioIds("")
	if err == nil {
		t.Error("did not error when passed empty code")
	}
}

func TestGetScenarioIds_ExpiredCode(t *testing.T) {
	_, err := accessCodeClient.GetScenarioIds(accessCodes[0] + "_expired")
	if err == nil {
		t.Error("did not error when passed expired access code")
	}
}

func TestGetCourseIds(t *testing.T) {
	ids, err := accessCodeClient.GetCourseIds(accessCodes[0])
	if err != nil {
		t.Error(err)
	}

	if len(ids) < 1 {
		t.Error("no course ids returned for valid access code")
	}

	var valid = true
	for i, v := range ids {
		if courses[i] != v {
			valid = false
		}
	}

	if !valid {
		t.Errorf("invalid courses returned, expected %s, got %s", courses, ids)
	}
}

func TestGetCourseIds_EmptyCode(t *testing.T) {
	_, err := accessCodeClient.GetCourseIds("")
	if err == nil {
		t.Error("did not error when passed empty access code")
	}
}

func TestGetCourseIds_ExpiredCode(t *testing.T) {
	_, err := accessCodeClient.GetCourseIds(accessCodes[0] + "_expired")
	if err == nil {
		t.Error("did not error when passed expired access code")
	}
}

func TestGetClosestAccessCode_UsingScenario(t *testing.T) {
	ac, err := accessCodeClient.GetClosestAccessCode("user", scenarios[0])
	if err != nil {
		t.Error(err)
	}

	// this access code should be the soonest-expiring code, which should be accessCodes[0]
	if ac != accessCodes[0] {
		t.Errorf("did not receive oldest access code, expected %s, received %s", accessCodes[0], ac)
	}
}

func TestGetClosestAccesscode_UsingCourse(t *testing.T) {
	ac, err := accessCodeClient.GetClosestAccessCode("user", courses[0])
	if err != nil {
		t.Error(err)
	}

	// this access code should be the soonest-expiring code, which should be accessCodes[0]
	if ac != accessCodes[0] {
		t.Errorf("did not receive oldest access code, expected %s, received %s", accessCodes[0], ac)
	}
}

func TestGetClosestAccessCode_BadUser(t *testing.T) {
	_, err := accessCodeClient.GetClosestAccessCode("baduser", scenarios[0])
	// this should return error as no user with "baduser" id should exist
	if err == nil {
		t.Error("did not error when passed nonexistent user id")
	}
}

func TestGetClosestAccessCode_ExpiredUser(t *testing.T) {
	_, err := accessCodeClient.GetClosestAccessCode("user_expired", scenarios[0])
	// this should return error as the user has no valid access codes (all expired)
	if err == nil {
		t.Error("did not error when searching for access codes that have expired")
	}
}

func createAccessCode(name string, code string, scenarios []string, courses []string, expiration string, vmSets []string, restrictedBind bool, rbValue string) hobbyfarmiov1.AccessCode {
	return hobbyfarmiov1.AccessCode{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: hobbyfarmiov1.AccessCodeSpec{
			Code:                code,
			Description:         code,
			Scenarios:           scenarios,
			Courses:             courses,
			Expiration:          expiration,
			VirtualMachineSets:  vmSets,
			RestrictedBind:      restrictedBind,
			RestrictedBindValue: rbValue,
		},
	}
}

func createUser(name string, id string, email string, password string, accessCodes []string, admin bool) hobbyfarmiov1.User {
	return hobbyfarmiov1.User{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec:       hobbyfarmiov1.UserSpec{
			Id: id,
			Email: email,
			Password: password, // doesn't matter since we're not logging in with this user
			AccessCodes: accessCodes,
			Admin: admin,
		},
	}
}