Hướng dẫn cấu hình Gin  với wire(DI)


1. Cài đặt 

    ```bash
    go install github.com/google/wire/cmd/wire@latest
    ```

    chạy: ```go mod tidy``` để cho nó cài  đặt  kiểu  ```npm i ``` bằng không nó báo lỗi rằng package này chưa có ^^

2. Phần cấu  hình 

    2.1 Trong file cần cấu hình

    - VD: đây là file user service

    ````go
    type UserController struct {
	    userSrv *services.UserService
    }

    func NewUserController(userSrv *services.UserService) *UserController {
        return &UserController{userSrv: userSrv}
    }
    ````

    - func NewUserController(userSrv *services.UserService) 
        -  bạn có  để ý  thấy rằng trong parameter của function có 1 biến không? đây  chính xác trong golang là  contructor. vậy thì giống như ngôn ngữ khác thì đây chính là kiểu contructor class base. đúng là khắm lọ. Mình cũng không hiểu sao nó không cho phép viết OOP nhưng nó lại có function contructor giống JS es5 ??? :))))
         Nếu không có thì WIRE  nó cũng báo lỗi  là không được inject và báo lỗi ngay trong phần khai báo file wire.build()


    - Những thứ khác như struct thì hãy dùng con trỏ (Pointer).

    2.2 Giờ là màn kịch hay
     -  tạo 1 file  go với tên tuỳ ý. vd: wire.go

    ```go
        // go:build wireinject
        //go:build wireinject
        // +build wireinject

        package di_wire

        import (
            "test/controllers"
            "test/repo"
            "test/services"

            "github.com/google/wire"
        )

        func InitUserRouterHandler() *controllers.UserController {
            panic(wire.Build(
                repo.NewUserRepo,
                services.NewUserService,
                controllers.NewUserController,
            ))
        }
    ```

    - 3 cái comment đầu không phải để cho vui, cũng ko phải thích viết gì viết mà nó phải có  3 dòng này. bằng không  golang  nó  build tào  lao =))) nếu thiếu dòng cuối thì  nó build xong báo duplicate function ngay lập tức. Lý do thì mình cũng chưa tìm hiểu nhưng đại loại vô tình chatgpt nó generate  ra 3 dòng đó và work ngay lập tức =)) Nếu không tin các hạ xoá file wire_gen.go và build thử là biết.  


    2.3 cuối cùng là chạy lệnh
    
    ```bash
        wire path_your_wire_file
    ```

    - nếu thành công thì nó sẽ chỉ ra đường dẫn file. nếu không nó sẽ báo lỗi ở wire.build ở DI nào. bạn buộc phải sửa trước khi build lại



    