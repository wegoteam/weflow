namespace go common

struct Response {
  1: required string msg
  2: required i32 code
  3: required bool success
  4: optional binary data
}

typedef map<string, string> Data