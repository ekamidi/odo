// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/asset.proto

package resources // import "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import common "google.golang.org/genproto/googleapis/ads/googleads/v1/common"
import enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Asset is a part of an ad which can be shared across multiple ads.
// It can be an image (ImageAsset), a video (YoutubeVideoAsset), etc.
type Asset struct {
	// The resource name of the asset.
	// Asset resource names have the form:
	//
	// `customers/{customer_id}/assets/{asset_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the asset.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Optional name of the asset.
	Name *wrappers.StringValue `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Type of the asset.
	Type enums.AssetTypeEnum_AssetType `protobuf:"varint,4,opt,name=type,proto3,enum=google.ads.googleads.v1.enums.AssetTypeEnum_AssetType" json:"type,omitempty"`
	// The specific type of the asset.
	//
	// Types that are valid to be assigned to AssetData:
	//	*Asset_YoutubeVideoAsset
	//	*Asset_MediaBundleAsset
	//	*Asset_ImageAsset
	AssetData            isAsset_AssetData `protobuf_oneof:"asset_data"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Asset) Reset()         { *m = Asset{} }
func (m *Asset) String() string { return proto.CompactTextString(m) }
func (*Asset) ProtoMessage()    {}
func (*Asset) Descriptor() ([]byte, []int) {
	return fileDescriptor_asset_c59e4d590e033b73, []int{0}
}
func (m *Asset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Asset.Unmarshal(m, b)
}
func (m *Asset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Asset.Marshal(b, m, deterministic)
}
func (dst *Asset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Asset.Merge(dst, src)
}
func (m *Asset) XXX_Size() int {
	return xxx_messageInfo_Asset.Size(m)
}
func (m *Asset) XXX_DiscardUnknown() {
	xxx_messageInfo_Asset.DiscardUnknown(m)
}

var xxx_messageInfo_Asset proto.InternalMessageInfo

func (m *Asset) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *Asset) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Asset) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Asset) GetType() enums.AssetTypeEnum_AssetType {
	if m != nil {
		return m.Type
	}
	return enums.AssetTypeEnum_UNSPECIFIED
}

type isAsset_AssetData interface {
	isAsset_AssetData()
}

type Asset_YoutubeVideoAsset struct {
	YoutubeVideoAsset *common.YoutubeVideoAsset `protobuf:"bytes,5,opt,name=youtube_video_asset,json=youtubeVideoAsset,proto3,oneof"`
}

type Asset_MediaBundleAsset struct {
	MediaBundleAsset *common.MediaBundleAsset `protobuf:"bytes,6,opt,name=media_bundle_asset,json=mediaBundleAsset,proto3,oneof"`
}

type Asset_ImageAsset struct {
	ImageAsset *common.ImageAsset `protobuf:"bytes,7,opt,name=image_asset,json=imageAsset,proto3,oneof"`
}

func (*Asset_YoutubeVideoAsset) isAsset_AssetData() {}

func (*Asset_MediaBundleAsset) isAsset_AssetData() {}

func (*Asset_ImageAsset) isAsset_AssetData() {}

func (m *Asset) GetAssetData() isAsset_AssetData {
	if m != nil {
		return m.AssetData
	}
	return nil
}

func (m *Asset) GetYoutubeVideoAsset() *common.YoutubeVideoAsset {
	if x, ok := m.GetAssetData().(*Asset_YoutubeVideoAsset); ok {
		return x.YoutubeVideoAsset
	}
	return nil
}

func (m *Asset) GetMediaBundleAsset() *common.MediaBundleAsset {
	if x, ok := m.GetAssetData().(*Asset_MediaBundleAsset); ok {
		return x.MediaBundleAsset
	}
	return nil
}

func (m *Asset) GetImageAsset() *common.ImageAsset {
	if x, ok := m.GetAssetData().(*Asset_ImageAsset); ok {
		return x.ImageAsset
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Asset) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Asset_OneofMarshaler, _Asset_OneofUnmarshaler, _Asset_OneofSizer, []interface{}{
		(*Asset_YoutubeVideoAsset)(nil),
		(*Asset_MediaBundleAsset)(nil),
		(*Asset_ImageAsset)(nil),
	}
}

func _Asset_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Asset)
	// asset_data
	switch x := m.AssetData.(type) {
	case *Asset_YoutubeVideoAsset:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.YoutubeVideoAsset); err != nil {
			return err
		}
	case *Asset_MediaBundleAsset:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.MediaBundleAsset); err != nil {
			return err
		}
	case *Asset_ImageAsset:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ImageAsset); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Asset.AssetData has unexpected type %T", x)
	}
	return nil
}

func _Asset_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Asset)
	switch tag {
	case 5: // asset_data.youtube_video_asset
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.YoutubeVideoAsset)
		err := b.DecodeMessage(msg)
		m.AssetData = &Asset_YoutubeVideoAsset{msg}
		return true, err
	case 6: // asset_data.media_bundle_asset
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.MediaBundleAsset)
		err := b.DecodeMessage(msg)
		m.AssetData = &Asset_MediaBundleAsset{msg}
		return true, err
	case 7: // asset_data.image_asset
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.ImageAsset)
		err := b.DecodeMessage(msg)
		m.AssetData = &Asset_ImageAsset{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Asset_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Asset)
	// asset_data
	switch x := m.AssetData.(type) {
	case *Asset_YoutubeVideoAsset:
		s := proto.Size(x.YoutubeVideoAsset)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Asset_MediaBundleAsset:
		s := proto.Size(x.MediaBundleAsset)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Asset_ImageAsset:
		s := proto.Size(x.ImageAsset)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Asset)(nil), "google.ads.googleads.v1.resources.Asset")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/asset.proto", fileDescriptor_asset_c59e4d590e033b73)
}

var fileDescriptor_asset_c59e4d590e033b73 = []byte{
	// 487 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0x49, 0xda, 0x0d, 0xe1, 0x0d, 0x04, 0xe6, 0x52, 0x8d, 0x09, 0x75, 0xa0, 0x49, 0x15,
	0x08, 0xa7, 0x19, 0x68, 0x87, 0x70, 0x4a, 0x25, 0x34, 0x36, 0x69, 0x68, 0x0a, 0x28, 0x12, 0xa8,
	0x52, 0x70, 0x6b, 0x13, 0x59, 0xaa, 0xed, 0x28, 0x76, 0x8a, 0xfa, 0x3a, 0x1c, 0x79, 0x14, 0x1e,
	0x85, 0x17, 0xe0, 0xc0, 0x05, 0xc5, 0x8e, 0x0d, 0x1a, 0xca, 0x7a, 0xfb, 0xec, 0xfc, 0xfe, 0xbf,
	0xef, 0xab, 0xed, 0x82, 0x17, 0xa5, 0x94, 0xe5, 0x8a, 0x46, 0x98, 0xa8, 0xc8, 0x96, 0x6d, 0xb5,
	0x8e, 0xa3, 0x9a, 0x2a, 0xd9, 0xd4, 0x4b, 0xaa, 0x22, 0xac, 0x14, 0xd5, 0xa8, 0xaa, 0xa5, 0x96,
	0xf0, 0xc8, 0x32, 0x08, 0x13, 0x85, 0x3c, 0x8e, 0xd6, 0x31, 0xf2, 0xf8, 0xc1, 0xb4, 0xcf, 0xb8,
	0x94, 0x9c, 0x4b, 0x61, 0x75, 0x85, 0xde, 0x54, 0x54, 0x59, 0xe9, 0x01, 0xea, 0x4b, 0x50, 0xd1,
	0x70, 0xf5, 0x4f, 0xa0, 0xe3, 0x1f, 0x77, 0xbc, 0x59, 0x2d, 0x9a, 0x2f, 0xd1, 0xd7, 0x1a, 0x57,
	0x15, 0xad, 0x9d, 0xef, 0xd0, 0xf9, 0x2a, 0x16, 0x61, 0x21, 0xa4, 0xc6, 0x9a, 0x49, 0xd1, 0x7d,
	0x7d, 0xf2, 0x7b, 0x00, 0x76, 0xd2, 0x56, 0x09, 0x9f, 0x82, 0xbb, 0x6e, 0xec, 0x42, 0x60, 0x4e,
	0x47, 0xc1, 0x38, 0x98, 0xdc, 0xc9, 0xf6, 0xdd, 0xe6, 0x3b, 0xcc, 0x29, 0x7c, 0x0e, 0x42, 0x46,
	0x46, 0xe1, 0x38, 0x98, 0xec, 0x9d, 0x3c, 0xea, 0x26, 0x45, 0xae, 0x33, 0x3a, 0x17, 0xfa, 0xf4,
	0x55, 0x8e, 0x57, 0x0d, 0xcd, 0x42, 0x46, 0xe0, 0x14, 0x0c, 0x8d, 0x68, 0x60, 0xf0, 0xc3, 0xff,
	0xf0, 0xf7, 0xba, 0x66, 0xa2, 0xb4, 0xbc, 0x21, 0xe1, 0x05, 0x18, 0xb6, 0xbf, 0x6c, 0x34, 0x1c,
	0x07, 0x93, 0x7b, 0x27, 0xa7, 0xa8, 0xef, 0x7c, 0xcd, 0x51, 0x20, 0x33, 0xf7, 0x87, 0x4d, 0x45,
	0xdf, 0x88, 0x86, 0xff, 0x5d, 0x65, 0xc6, 0x01, 0x97, 0xe0, 0xe1, 0x46, 0x36, 0xba, 0x59, 0xd0,
	0x62, 0xcd, 0x08, 0x95, 0x85, 0x39, 0xb9, 0xd1, 0x8e, 0x19, 0x26, 0xee, 0x55, 0xdb, 0x7b, 0x41,
	0x1f, 0x6d, 0x34, 0x6f, 0x93, 0xc6, 0xfc, 0xf6, 0x56, 0xf6, 0x60, 0x73, 0x7d, 0x13, 0x7e, 0x06,
	0x90, 0x53, 0xc2, 0x70, 0xb1, 0x68, 0x04, 0x59, 0xd1, 0xae, 0xc7, 0xae, 0xe9, 0x31, 0xdd, 0xd6,
	0xe3, 0xb2, 0x4d, 0xce, 0x4c, 0xd0, 0xb5, 0xb8, 0xcf, 0xaf, 0xed, 0xc1, 0x4b, 0xb0, 0xc7, 0x38,
	0x2e, 0x9d, 0xfa, 0xb6, 0x51, 0x3f, 0xdb, 0xa6, 0x3e, 0x6f, 0x23, 0x4e, 0x0a, 0x98, 0x5f, 0xcd,
	0xf6, 0x01, 0xb0, 0x2f, 0x88, 0x60, 0x8d, 0x67, 0xbf, 0x02, 0x70, 0xbc, 0x94, 0x1c, 0x6d, 0x7d,
	0xc7, 0x33, 0x60, 0xe2, 0x57, 0xed, 0xd5, 0x5d, 0x05, 0x9f, 0x2e, 0xba, 0x40, 0x29, 0x57, 0x58,
	0x94, 0x48, 0xd6, 0x65, 0x54, 0x52, 0x61, 0x2e, 0xd6, 0xbd, 0xd9, 0x8a, 0xa9, 0x1b, 0xfe, 0x46,
	0xaf, 0x7d, 0xf5, 0x2d, 0x1c, 0x9c, 0xa5, 0xe9, 0xf7, 0xf0, 0xe8, 0xcc, 0x2a, 0x53, 0xa2, 0x90,
	0x2d, 0xdb, 0x2a, 0x8f, 0x51, 0xe6, 0xc8, 0x1f, 0x8e, 0x99, 0xa7, 0x44, 0xcd, 0x3d, 0x33, 0xcf,
	0xe3, 0xb9, 0x67, 0x7e, 0x86, 0xc7, 0xf6, 0x43, 0x92, 0xa4, 0x44, 0x25, 0x89, 0xa7, 0x92, 0x24,
	0x8f, 0x93, 0xc4, 0x73, 0x8b, 0x5d, 0x33, 0xec, 0xcb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x46,
	0x7a, 0xa1, 0x74, 0xf2, 0x03, 0x00, 0x00,
}
