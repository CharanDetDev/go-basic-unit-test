# The Go Language Programming : Basic Unit Testing

> ## **vscode Setting**

- Monitor viewer On vscode :
  แก้ไขเพิ่มเติมที่ setting.json :

```json
{
  "go.coverOnSave": true,
  "go.coverOnSingleTest": true,
  "go.coverageDecorator": {
    "type": "gutter",
    "coveredHighlightColor": "rgba(64,128,128,0.5)",
    "uncoveredHighlightColor": "rgba(128,64,64,0.25)",
    "coveredGutterStyle": "blockgreen",
    "uncoveredGutterStyle": "blockred"
  }
}
```

- หลักการ AAA ของการทำ Unit test
  - Arrage : คือการเตรียมข้อมูลสำหรับการทำ Unit test
  - Action : คือการเรียกใช้ Function ที่ต้องการทดสอบ
  - Asset : คือการตรวจสอบผลลัพธ์ ว่าเป็นไปตามที่คาดหวังหรือไม่

> ## **Example**
>
> _Basic Unit Testing_

```golang

package main

import "testing"


func main() {
	defer database.ConnectionClose()

	hello := "Hello"
	fmt.Println(hello, HelloWorld(hello))

	app := fiber.New()
	router := route.NewRoute()
	router.InitRoute(app)

	app.Listen(config.Env.API_PORT)
}

func HelloWorld(hello string) string {
	if hello != "" {
		return "World"
	}

	return ""
}

func TestHelloWorld(t *testing.T) { //! Function TestHelloWorld จะถูกสร้างขึ้นมาอัตโนมัติ

	//* Arrage
	type args struct {
		hello string //! argument ที่ต้องการส่งไปทดสอบ
	}
	tests := []struct {
		name string //! ชื่อของ test case
		args args   //! explicit field ของ argument
		want string //! expected
	}{
		// TODO: Add test cases.
		{
			name: "Case_Hello_World",
			args: args{
				hello: "Hello",
			},
			want: "World",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			//* Action
			got := HelloWorld(tt.args.hello)

			//* Assert
			if got != tt.want {
				t.Errorf("HelloWorld() = %v, want %v", got, tt.want)
			}
		})
	}
}

```
