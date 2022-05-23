// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: proto/employee.proto

package protobuf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Occupation int32

const (
	Occupation_UNKNOWN  Occupation = 0
	Occupation_ENGINEER Occupation = 1
	Occupation_DESIGNER Occupation = 2
	Occupation_MANAGER  Occupation = 3
)

// Enum value maps for Occupation.
var (
	Occupation_name = map[int32]string{
		0: "UNKNOWN",
		1: "ENGINEER",
		2: "DESIGNER",
		3: "MANAGER",
	}
	Occupation_value = map[string]int32{
		"UNKNOWN":  0,
		"ENGINEER": 1,
		"DESIGNER": 2,
		"MANAGER":  3,
	}
)

func (x Occupation) Enum() *Occupation {
	p := new(Occupation)
	*p = x
	return p
}

func (x Occupation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Occupation) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_employee_proto_enumTypes[0].Descriptor()
}

func (Occupation) Type() protoreflect.EnumType {
	return &file_proto_employee_proto_enumTypes[0]
}

func (x Occupation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Occupation.Descriptor instead.
func (Occupation) EnumDescriptor() ([]byte, []int) {
	return file_proto_employee_proto_rawDescGZIP(), []int{0}
}

type Employee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                int32                       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name              string                      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email             string                      `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Occupation        Occupation                  `protobuf:"varint,4,opt,name=occupation,proto3,enum=employee.Occupation" json:"occupation,omitempty"`
	ThirdPartyAccount []string                    `protobuf:"bytes,5,rep,name=third_party_account,json=thirdPartyAccount,proto3" json:"third_party_account,omitempty"`
	Products          map[string]*Company_Product `protobuf:"bytes,6,rep,name=products,proto3" json:"products,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Types that are assignable to Profile:
	//	*Employee_Text
	//	*Employee_Url
	Profile    isEmployee_Profile `protobuf_oneof:"profile"`
	JoinedDate *Date              `protobuf:"bytes,9,opt,name=joinedDate,proto3" json:"joinedDate,omitempty"`
}

func (x *Employee) Reset() {
	*x = Employee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_employee_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Employee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Employee) ProtoMessage() {}

func (x *Employee) ProtoReflect() protoreflect.Message {
	mi := &file_proto_employee_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Employee.ProtoReflect.Descriptor instead.
func (*Employee) Descriptor() ([]byte, []int) {
	return file_proto_employee_proto_rawDescGZIP(), []int{0}
}

func (x *Employee) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Employee) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Employee) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Employee) GetOccupation() Occupation {
	if x != nil {
		return x.Occupation
	}
	return Occupation_UNKNOWN
}

func (x *Employee) GetThirdPartyAccount() []string {
	if x != nil {
		return x.ThirdPartyAccount
	}
	return nil
}

func (x *Employee) GetProducts() map[string]*Company_Product {
	if x != nil {
		return x.Products
	}
	return nil
}

func (m *Employee) GetProfile() isEmployee_Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

func (x *Employee) GetText() string {
	if x, ok := x.GetProfile().(*Employee_Text); ok {
		return x.Text
	}
	return ""
}

func (x *Employee) GetUrl() *URL {
	if x, ok := x.GetProfile().(*Employee_Url); ok {
		return x.Url
	}
	return nil
}

func (x *Employee) GetJoinedDate() *Date {
	if x != nil {
		return x.JoinedDate
	}
	return nil
}

type isEmployee_Profile interface {
	isEmployee_Profile()
}

type Employee_Text struct {
	Text string `protobuf:"bytes,7,opt,name=text,proto3,oneof"`
}

type Employee_Url struct {
	Url *URL `protobuf:"bytes,8,opt,name=url,proto3,oneof"`
}

func (*Employee_Text) isEmployee_Profile() {}

func (*Employee_Url) isEmployee_Profile() {}

type Company struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Company) Reset() {
	*x = Company{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_employee_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Company) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Company) ProtoMessage() {}

func (x *Company) ProtoReflect() protoreflect.Message {
	mi := &file_proto_employee_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Company.ProtoReflect.Descriptor instead.
func (*Company) Descriptor() ([]byte, []int) {
	return file_proto_employee_proto_rawDescGZIP(), []int{1}
}

type URL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *URL) Reset() {
	*x = URL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_employee_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *URL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*URL) ProtoMessage() {}

func (x *URL) ProtoReflect() protoreflect.Message {
	mi := &file_proto_employee_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use URL.ProtoReflect.Descriptor instead.
func (*URL) Descriptor() ([]byte, []int) {
	return file_proto_employee_proto_rawDescGZIP(), []int{2}
}

type Company_Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Company_Product) Reset() {
	*x = Company_Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_employee_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Company_Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Company_Product) ProtoMessage() {}

func (x *Company_Product) ProtoReflect() protoreflect.Message {
	mi := &file_proto_employee_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Company_Product.ProtoReflect.Descriptor instead.
func (*Company_Product) Descriptor() ([]byte, []int) {
	return file_proto_employee_proto_rawDescGZIP(), []int{1, 0}
}

var File_proto_employee_proto protoreflect.FileDescriptor

var file_proto_employee_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x1a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb0, 0x03, 0x0a, 0x08, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x34, 0x0a, 0x0a, 0x6f, 0x63, 0x63,
	0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e,
	0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x6f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2e, 0x0a, 0x13, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x74, 0x68,
	0x69, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x3c, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x45, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x14, 0x0a,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x12, 0x21, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x55, 0x52, 0x4c, 0x48,
	0x00, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x2a, 0x0a, 0x0a, 0x6a, 0x6f, 0x69, 0x6e, 0x65, 0x64,
	0x44, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x0a, 0x6a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x44, 0x61,
	0x74, 0x65, 0x1a, 0x56, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x14, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x1a, 0x09, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x05, 0x0a, 0x03, 0x55,
	0x52, 0x4c, 0x2a, 0x42, 0x0a, 0x0a, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0c, 0x0a,
	0x08, 0x45, 0x4e, 0x47, 0x49, 0x4e, 0x45, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x44,
	0x45, 0x53, 0x49, 0x47, 0x4e, 0x45, 0x52, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x41, 0x4e,
	0x41, 0x47, 0x45, 0x52, 0x10, 0x03, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_employee_proto_rawDescOnce sync.Once
	file_proto_employee_proto_rawDescData = file_proto_employee_proto_rawDesc
)

func file_proto_employee_proto_rawDescGZIP() []byte {
	file_proto_employee_proto_rawDescOnce.Do(func() {
		file_proto_employee_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_employee_proto_rawDescData)
	})
	return file_proto_employee_proto_rawDescData
}

var file_proto_employee_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_employee_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_employee_proto_goTypes = []interface{}{
	(Occupation)(0),         // 0: employee.Occupation
	(*Employee)(nil),        // 1: employee.Employee
	(*Company)(nil),         // 2: employee.Company
	(*URL)(nil),             // 3: employee.URL
	nil,                     // 4: employee.Employee.ProductsEntry
	(*Company_Product)(nil), // 5: employee.Company.Product
	(*Date)(nil),            // 6: date.Date
}
var file_proto_employee_proto_depIdxs = []int32{
	0, // 0: employee.Employee.occupation:type_name -> employee.Occupation
	4, // 1: employee.Employee.products:type_name -> employee.Employee.ProductsEntry
	3, // 2: employee.Employee.url:type_name -> employee.URL
	6, // 3: employee.Employee.joinedDate:type_name -> date.Date
	5, // 4: employee.Employee.ProductsEntry.value:type_name -> employee.Company.Product
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_employee_proto_init() }
func file_proto_employee_proto_init() {
	if File_proto_employee_proto != nil {
		return
	}
	file_proto_date_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_employee_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Employee); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_employee_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Company); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_employee_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*URL); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_employee_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Company_Product); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_proto_employee_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Employee_Text)(nil),
		(*Employee_Url)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_employee_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_employee_proto_goTypes,
		DependencyIndexes: file_proto_employee_proto_depIdxs,
		EnumInfos:         file_proto_employee_proto_enumTypes,
		MessageInfos:      file_proto_employee_proto_msgTypes,
	}.Build()
	File_proto_employee_proto = out.File
	file_proto_employee_proto_rawDesc = nil
	file_proto_employee_proto_goTypes = nil
	file_proto_employee_proto_depIdxs = nil
}
