// 원글 링크 : https://code.google.com/p/golang-korea/wiki/EffectiveGo#명칭(Names)
/*

Package names
package가 import될 때, package 이름은 package 내용에 대한 접근자가 됩니다.

import "bytes"

위와 같이 썼다면 importing된 package는 byte에 관해 다룰 것입니다.
만약 모든 사람들이 같은 package를 참조하기 위해 같은 이름을 사용한다면 편리할 것이며,
그것은 곧 package의 이름이 반드시 좋아야 한다는 것을 의미합니다.: 짧고, 간결해야하며,
유추가능해야 함. 규정에 따라, package 이름은 소문자를 사용하며, 한단어 이름 을 사용합니다.

밑줄(_) 또는 대소문자를 섞어서 사용할 필요가 없습니다. 당신의 package를 사용하는 모든 사람이
그 이름을 사용할 것이기 때문에 간결함이 문제가 될 수도 있지만, package 이름은 import를 위한
기본 이름으로만 사용되기 때문에 중복 및 충돌에 대해서는 걱정할 필요가 없습니다.

(모든 소스코드에 있어서 유일한 이름일 필요는 없으며, 가끔 충돌이 발생하는 경우에 한해서

국지적으로(local) 쓰기위해 다른 이름을 사용하면 됩니다.) import되는 경우 파일의 이름으로

어떤 package가 사용되어 졌는지를 결정할 것 이기 때문에 혼동되는 경우는 거의 없습니다.

또 다른 규정은 package 이름은 package 소스 디렉토리에 기반한 이름이어야 한다는 것입니다.

즉, package 소스 디렉토리가 src/pkg/container/vector 라고 가정하면 그
 package는 "container/vector"로 import 되어지며 이름은 vector 입니다.
(container_vector 또는 containerVector가 아닙니다.)

package를 import하려면 package 내용을 참조하기 위해 package의 이름을 사용할 것입니다.
'.' 기호는 test나 일반적이지 않은 상황을 위한 것이며 꼭 필요하지 않다면 피해야합니다.

예를 들어 bufio package에서 buffered reader type은 BufReader?가 아니라 Reader로 불립니다.
왜냐하면 사용자들은 buffered reader type을 bufio.Reader로써 볼 것이고,
 그것이 분명하고 간결한 이름이기 때문이며, 더욱이, import된 객체들은 항상
그 객체의 package 이름과 함께 사용되기 때문입니다. bufio.Reader는 io.Reader와 충돌을 일으키지 않습니다.

유사한 예로, ring.Ring(Go언어에서 한 생성자의 정의임)의 새로운 개체를 만들기 위한 함수는 일반적으로 NewRing?으로 호출될 것 입니다.
그렇지만 Ring은 ring package에 의해 export된 유일한 type이며,
그 package는 ring이기 때문에 새로운 개체는 그냥 New라고 호출되며 그 package의 client들에게는 ring.New로 보일 것입니다.
좋은 이름을 선택하기 위해서는 package 구조를 사용하십시요.

또 다른 간단한 예는 once.Do 입니다. once.Do(setup)은 읽기에 쉽고
once.DoOrWaitUntilDone?(setup) 보다 좋습니다. 긴 이름들은 읽기에 쉽지 않습니다.
 만약에 그 이름이 뭔가 복잡하고 어려운 무언가를 대표한다면,
그 이름에 모든 정보를 넣기보다는 doc 주석으로 남기는 것이 더 좋습니다.

*/