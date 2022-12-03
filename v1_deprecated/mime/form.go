// v1 of Waggy has been deprecated. Please use v2 by running the following
// command: go get github.com/syke99/waggy/v2

package mime

//
//type Form struct {
//	values map[string][]string
//}
//
//// GetForm initializes a *Form
//func GetForm() *Form {
//	f := Form{values: make(map[string][]string)}
//
//	return &f
//}
//
//// Add works just like URL.Query.Add() in net/url
//func (f *Form) Add(key string, value string) {
//	f.values[key] = append(f.values[key], value)
//}
//
//// Get works just like URL.Query.Get() in net/url
//func (f *Form) Get(key string) string {
//	if _, ok := f.values[key]; !ok {
//		return ""
//	}
//
//	return f.values[key][0]
//}
//
//// Has works just like URL.Query.Has() in net/url
//func (f *Form) Has(key string) bool {
//	if _, ok := f.values[key]; !ok {
//		return false
//	}
//
//	return true
//}
//
//// Set works just like URL.Query.Set() in net/url
//func (f *Form) Set(key string, value string) {
//	f.Del(key)
//
//	f.Add(key, value)
//}
//
//// Del works just like URL.Query.Del() in net/url
//func (f *Form) Del(key string) {
//	delete(f.values, key)
//}
