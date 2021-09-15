

| Tag                  | Param       | Explanation                                                  |
| -------------------- | ----------- | ------------------------------------------------------------ |
| required             |             | field is required                                            |
|                      |             |                                                              |
| required_if          | A foo       | field is required if A == "foo"                              |
| required_if          | A foo B bar | field is required if A == "foo" and B == "bar"               |
|                      |             |                                                              |
| required_unless      | A foo       | field is required if A != "foo" (field is required unless A == "foo") |
| required_unless      | A foo B bar | field is required if A != "foo" and B != "bar" (field is required unless A == "foo" or B == "bar") |
|                      |             |                                                              |
| required_with        | A           | field is required if A == ∅                                  |
| required_with (BUG)  | A B         | field is required if A != ∅ or B != ∅                        |
| required_with_all    | A B         | field is required if A != ∅ and B != ∅                       |
|                      |             |                                                              |
| required_without     | A           | field is required if A == ∅                                  |
| required_without     | A B         | field is required if A == ∅ or B == ∅                        |
| required_without_all | A B         | field is required if A == ∅ and B == ∅                       |

∅ represents zero value


