// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	gits "github.com/jenkins-x/jx/pkg/gits"
)

func AnyPtrToGitsGitPullRequest() *gits.GitPullRequest {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*gits.GitPullRequest))(nil)).Elem()))
	var nullValue *gits.GitPullRequest
	return nullValue
}

func EqPtrToGitsGitPullRequest(value *gits.GitPullRequest) *gits.GitPullRequest {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *gits.GitPullRequest
	return nullValue
}
