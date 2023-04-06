namespace go flow
include "common.thrift"

struct Person {
  1: required i32 id
  2: optional string name
  3: optional i32 age
}

service FlowService {
  list<Person> getList(),
  map<string, i32> getMap(1: string key),
  common.Response getPerson(1: string name)
}
