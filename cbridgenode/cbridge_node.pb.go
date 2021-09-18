// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: cbridge_node.proto

package cbridgenode

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

type TransferStatus int32

const (
	// pending -> locked -> confirmed
	//               |----> refunded
	TransferStatus_TRANSFER_STATUS_UNDEFINED           TransferStatus = 0
	TransferStatus_TRANSFER_STATUS_LOCKED              TransferStatus = 1
	TransferStatus_TRANSFER_STATUS_CONFIRMED           TransferStatus = 2
	TransferStatus_TRANSFER_STATUS_REFUNDED            TransferStatus = 3
	TransferStatus_TRANSFER_STATUS_TRANSFER_IN_START   TransferStatus = 4
	TransferStatus_TRANSFER_STATUS_TRANSFER_IN_PENDING TransferStatus = 5
	TransferStatus_TRANSFER_STATUS_CONFIRM_PENDING     TransferStatus = 6
	TransferStatus_TRANSFER_STATUS_REFUND_PENDING      TransferStatus = 7
)

// Enum value maps for TransferStatus.
var (
	TransferStatus_name = map[int32]string{
		0: "TRANSFER_STATUS_UNDEFINED",
		1: "TRANSFER_STATUS_LOCKED",
		2: "TRANSFER_STATUS_CONFIRMED",
		3: "TRANSFER_STATUS_REFUNDED",
		4: "TRANSFER_STATUS_TRANSFER_IN_START",
		5: "TRANSFER_STATUS_TRANSFER_IN_PENDING",
		6: "TRANSFER_STATUS_CONFIRM_PENDING",
		7: "TRANSFER_STATUS_REFUND_PENDING",
	}
	TransferStatus_value = map[string]int32{
		"TRANSFER_STATUS_UNDEFINED":           0,
		"TRANSFER_STATUS_LOCKED":              1,
		"TRANSFER_STATUS_CONFIRMED":           2,
		"TRANSFER_STATUS_REFUNDED":            3,
		"TRANSFER_STATUS_TRANSFER_IN_START":   4,
		"TRANSFER_STATUS_TRANSFER_IN_PENDING": 5,
		"TRANSFER_STATUS_CONFIRM_PENDING":     6,
		"TRANSFER_STATUS_REFUND_PENDING":      7,
	}
)

func (x TransferStatus) Enum() *TransferStatus {
	p := new(TransferStatus)
	*p = x
	return p
}

func (x TransferStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransferStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_cbridge_node_proto_enumTypes[0].Descriptor()
}

func (TransferStatus) Type() protoreflect.EnumType {
	return &file_cbridge_node_proto_enumTypes[0]
}

func (x TransferStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransferStatus.Descriptor instead.
func (TransferStatus) EnumDescriptor() ([]byte, []int) {
	return file_cbridge_node_proto_rawDescGZIP(), []int{0}
}

type TransferType int32

const (
	TransferType_TRANSFER_TYPE_UNDEFINED TransferType = 0
	TransferType_TRANSFER_TYPE_OUT       TransferType = 1
	TransferType_TRANSFER_TYPE_IN        TransferType = 2
)

// Enum value maps for TransferType.
var (
	TransferType_name = map[int32]string{
		0: "TRANSFER_TYPE_UNDEFINED",
		1: "TRANSFER_TYPE_OUT",
		2: "TRANSFER_TYPE_IN",
	}
	TransferType_value = map[string]int32{
		"TRANSFER_TYPE_UNDEFINED": 0,
		"TRANSFER_TYPE_OUT":       1,
		"TRANSFER_TYPE_IN":        2,
	}
)

func (x TransferType) Enum() *TransferType {
	p := new(TransferType)
	*p = x
	return p
}

func (x TransferType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransferType) Descriptor() protoreflect.EnumDescriptor {
	return file_cbridge_node_proto_enumTypes[1].Descriptor()
}

func (TransferType) Type() protoreflect.EnumType {
	return &file_cbridge_node_proto_enumTypes[1]
}

func (x TransferType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransferType.Descriptor instead.
func (TransferType) EnumDescriptor() ([]byte, []int) {
	return file_cbridge_node_proto_rawDescGZIP(), []int{1}
}

type ErrorCode int32

const (
	ErrorCode_INVALID ErrorCode = 0
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0: "INVALID",
	}
	ErrorCode_value = map[string]int32{
		"INVALID": 0,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_cbridge_node_proto_enumTypes[2].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_cbridge_node_proto_enumTypes[2]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_cbridge_node_proto_rawDescGZIP(), []int{2}
}

type CBridgeConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RelayNodeName string         `protobuf:"bytes,1,opt,name=relay_node_name,json=relayNodeName,proto3" json:"relay_node_name,omitempty"`
	ChainConfig   []*ChainConfig `protobuf:"bytes,2,rep,name=chain_config,json=chainConfig,proto3" json:"chain_config,omitempty"`
	Db            string         `protobuf:"bytes,3,opt,name=db,proto3" json:"db,omitempty"` // host:port
	Gateway       string         `protobuf:"bytes,4,opt,name=gateway,proto3" json:"gateway,omitempty"`
}

func (x *CBridgeConfig) Reset() {
	*x = CBridgeConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cbridge_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CBridgeConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CBridgeConfig) ProtoMessage() {}

func (x *CBridgeConfig) ProtoReflect() protoreflect.Message {
	mi := &file_cbridge_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CBridgeConfig.ProtoReflect.Descriptor instead.
func (*CBridgeConfig) Descriptor() ([]byte, []int) {
	return file_cbridge_node_proto_rawDescGZIP(), []int{0}
}

func (x *CBridgeConfig) GetRelayNodeName() string {
	if x != nil {
		return x.RelayNodeName
	}
	return ""
}

func (x *CBridgeConfig) GetChainConfig() []*ChainConfig {
	if x != nil {
		return x.ChainConfig
	}
	return nil
}

func (x *CBridgeConfig) GetDb() string {
	if x != nil {
		return x.Db
	}
	return ""
}

func (x *CBridgeConfig) GetGateway() string {
	if x != nil {
		return x.Gateway
	}
	return ""
}

type ChainConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoint         string            `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	ChainId          uint64            `protobuf:"varint,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	ContractAddress  string            `protobuf:"bytes,3,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	TokenConfig      []*TokenConfig    `protobuf:"bytes,4,rep,name=token_config,json=tokenConfig,proto3" json:"token_config,omitempty"` // supported token address
	WatchConfig      *WatchConfig      `protobuf:"bytes,5,opt,name=watch_config,json=watchConfig,proto3" json:"watch_config,omitempty"`
	ForceGasGwei     uint64            `protobuf:"varint,6,opt,name=force_gas_gwei,json=forceGasGwei,proto3" json:"force_gas_gwei,omitempty"`
	FeeRate          uint64            `protobuf:"varint,7,opt,name=fee_rate,json=feeRate,proto3" json:"fee_rate,omitempty"` // ten-thousandth, if percentage_fee is 1, then the relay node will use 1/10000 of the transfer amount
	GasTokenName     string            `protobuf:"bytes,8,opt,name=gas_token_name,json=gasTokenName,proto3" json:"gas_token_name,omitempty"`
	GasTokenDecimal  uint64            `protobuf:"varint,9,opt,name=gas_token_decimal,json=gasTokenDecimal,proto3" json:"gas_token_decimal,omitempty"`
	TransactorConfig *TransactorConfig `protobuf:"bytes,10,opt,name=transactor_config,json=transactorConfig,proto3" json:"transactor_config,omitempty"`
}

func (x *ChainConfig) Reset() {
	*x = ChainConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cbridge_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChainConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChainConfig) ProtoMessage() {}

func (x *ChainConfig) ProtoReflect() protoreflect.Message {
	mi := &file_cbridge_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChainConfig.ProtoReflect.Descriptor instead.
func (*ChainConfig) Descriptor() ([]byte, []int) {
	return file_cbridge_node_proto_rawDescGZIP(), []int{1}
}

func (x *ChainConfig) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *ChainConfig) GetChainId() uint64 {
	if x != nil {
		return x.ChainId
	}
	return 0
}

func (x *ChainConfig) GetContractAddress() string {
	if x != nil {
		return x.ContractAddress
	}
	return ""
}

func (x *ChainConfig) GetTokenConfig() []*TokenConfig {
	if x != nil {
		return x.TokenConfig
	}
	return nil
}

func (x *ChainConfig) GetWatchConfig() *WatchConfig {
	if x != nil {
		return x.WatchConfig
	}
	return nil
}

func (x *ChainConfig) GetForceGasGwei() uint64 {
	if x != nil {
		return x.ForceGasGwei
	}
	return 0
}

func (x *ChainConfig) GetFeeRate() uint64 {
	if x != nil {
		return x.FeeRate
	}
	return 0
}

func (x *ChainConfig) GetGasTokenName() string {
	if x != nil {
		return x.GasTokenName
	}
	return ""
}

func (x *ChainConfig) GetGasTokenDecimal() uint64 {
	if x != nil {
		return x.GasTokenDecimal
	}
	return 0
}

func (x *ChainConfig) GetTransactorConfig() *TransactorConfig {
	if x != nil {
		return x.TransactorConfig
	}
	return nil
}

type TokenConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenName    string `protobuf:"bytes,1,opt,name=token_name,json=tokenName,proto3" json:"token_name,omitempty"`
	TokenAddress string `protobuf:"bytes,2,opt,name=token_address,json=tokenAddress,proto3" json:"token_address,omitempty"`
	TokenDecimal uint64 `protobuf:"varint,3,opt,name=token_decimal,json=tokenDecimal,proto3" json:"token_decimal,omitempty"`
}

func (x *TokenConfig) Reset() {
	*x = TokenConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cbridge_node_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenConfig) ProtoMessage() {}

func (x *TokenConfig) ProtoReflect() protoreflect.Message {
	mi := &file_cbridge_node_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenConfig.ProtoReflect.Descriptor instead.
func (*TokenConfig) Descriptor() ([]byte, []int) {
	return file_cbridge_node_proto_rawDescGZIP(), []int{2}
}

func (x *TokenConfig) GetTokenName() string {
	if x != nil {
		return x.TokenName
	}
	return ""
}

func (x *TokenConfig) GetTokenAddress() string {
	if x != nil {
		return x.TokenAddress
	}
	return ""
}

func (x *TokenConfig) GetTokenDecimal() uint64 {
	if x != nil {
		return x.TokenDecimal
	}
	return 0
}

type WatchConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PollingInterval   uint64 `protobuf:"varint,1,opt,name=polling_interval,json=pollingInterval,proto3" json:"polling_interval,omitempty"`
	BlockDelay        uint64 `protobuf:"varint,2,opt,name=block_delay,json=blockDelay,proto3" json:"block_delay,omitempty"`
	MaxBlockDelta     uint64 `protobuf:"varint,3,opt,name=max_block_delta,json=maxBlockDelta,proto3" json:"max_block_delta,omitempty"`
	ForwardBlockDelay uint64 `protobuf:"varint,4,opt,name=forward_block_delay,json=forwardBlockDelay,proto3" json:"forward_block_delay,omitempty"`
}

func (x *WatchConfig) Reset() {
	*x = WatchConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cbridge_node_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchConfig) ProtoMessage() {}

func (x *WatchConfig) ProtoReflect() protoreflect.Message {
	mi := &file_cbridge_node_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchConfig.ProtoReflect.Descriptor instead.
func (*WatchConfig) Descriptor() ([]byte, []int) {
	return file_cbridge_node_proto_rawDescGZIP(), []int{3}
}

func (x *WatchConfig) GetPollingInterval() uint64 {
	if x != nil {
		return x.PollingInterval
	}
	return 0
}

func (x *WatchConfig) GetBlockDelay() uint64 {
	if x != nil {
		return x.BlockDelay
	}
	return 0
}

func (x *WatchConfig) GetMaxBlockDelta() uint64 {
	if x != nil {
		return x.MaxBlockDelta
	}
	return 0
}

func (x *WatchConfig) GetForwardBlockDelay() uint64 {
	if x != nil {
		return x.ForwardBlockDelay
	}
	return 0
}

type TransactorConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GasLimit            uint64  `protobuf:"varint,1,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	AddGasGwei          uint64  `protobuf:"varint,2,opt,name=add_gas_gwei,json=addGasGwei,proto3" json:"add_gas_gwei,omitempty"`
	AddGasEstimateRatio float64 `protobuf:"fixed64,3,opt,name=add_gas_estimate_ratio,json=addGasEstimateRatio,proto3" json:"add_gas_estimate_ratio,omitempty"`
}

func (x *TransactorConfig) Reset() {
	*x = TransactorConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cbridge_node_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactorConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactorConfig) ProtoMessage() {}

func (x *TransactorConfig) ProtoReflect() protoreflect.Message {
	mi := &file_cbridge_node_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactorConfig.ProtoReflect.Descriptor instead.
func (*TransactorConfig) Descriptor() ([]byte, []int) {
	return file_cbridge_node_proto_rawDescGZIP(), []int{4}
}

func (x *TransactorConfig) GetGasLimit() uint64 {
	if x != nil {
		return x.GasLimit
	}
	return 0
}

func (x *TransactorConfig) GetAddGasGwei() uint64 {
	if x != nil {
		return x.AddGasGwei
	}
	return 0
}

func (x *TransactorConfig) GetAddGasEstimateRatio() float64 {
	if x != nil {
		return x.AddGasEstimateRatio
	}
	return 0
}

type Transfer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransferId   []byte         `protobuf:"bytes,1,opt,name=transfer_id,json=transferId,proto3" json:"transfer_id,omitempty"`
	CreateTs     uint64         `protobuf:"varint,2,opt,name=create_ts,json=createTs,proto3" json:"create_ts,omitempty"`
	Status       TransferStatus `protobuf:"varint,3,opt,name=status,proto3,enum=cbridgenode.TransferStatus" json:"status,omitempty"`
	Type         TransferType   `protobuf:"varint,4,opt,name=type,proto3,enum=cbridgenode.TransferType" json:"type,omitempty"`
	ContractAddr []byte         `protobuf:"bytes,5,opt,name=contract_addr,json=contractAddr,proto3" json:"contract_addr,omitempty"`
	Amount       uint64         `protobuf:"varint,6,opt,name=amount,proto3" json:"amount,omitempty"`
	Fee          uint64         `protobuf:"varint,7,opt,name=fee,proto3" json:"fee,omitempty"`
	ChainId      uint64         `protobuf:"varint,8,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	FromAddr     []byte         `protobuf:"bytes,9,opt,name=from_addr,json=fromAddr,proto3" json:"from_addr,omitempty"`
	ToAddr       []byte         `protobuf:"bytes,10,opt,name=to_addr,json=toAddr,proto3" json:"to_addr,omitempty"`
	TimeLock     uint64         `protobuf:"varint,11,opt,name=time_lock,json=timeLock,proto3" json:"time_lock,omitempty"`
}

func (x *Transfer) Reset() {
	*x = Transfer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cbridge_node_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transfer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transfer) ProtoMessage() {}

func (x *Transfer) ProtoReflect() protoreflect.Message {
	mi := &file_cbridge_node_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transfer.ProtoReflect.Descriptor instead.
func (*Transfer) Descriptor() ([]byte, []int) {
	return file_cbridge_node_proto_rawDescGZIP(), []int{5}
}

func (x *Transfer) GetTransferId() []byte {
	if x != nil {
		return x.TransferId
	}
	return nil
}

func (x *Transfer) GetCreateTs() uint64 {
	if x != nil {
		return x.CreateTs
	}
	return 0
}

func (x *Transfer) GetStatus() TransferStatus {
	if x != nil {
		return x.Status
	}
	return TransferStatus_TRANSFER_STATUS_UNDEFINED
}

func (x *Transfer) GetType() TransferType {
	if x != nil {
		return x.Type
	}
	return TransferType_TRANSFER_TYPE_UNDEFINED
}

func (x *Transfer) GetContractAddr() []byte {
	if x != nil {
		return x.ContractAddr
	}
	return nil
}

func (x *Transfer) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Transfer) GetFee() uint64 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (x *Transfer) GetChainId() uint64 {
	if x != nil {
		return x.ChainId
	}
	return 0
}

func (x *Transfer) GetFromAddr() []byte {
	if x != nil {
		return x.FromAddr
	}
	return nil
}

func (x *Transfer) GetToAddr() []byte {
	if x != nil {
		return x.ToAddr
	}
	return nil
}

func (x *Transfer) GetTimeLock() uint64 {
	if x != nil {
		return x.TimeLock
	}
	return 0
}

type CbridgeNodeError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code ErrorCode `protobuf:"varint,1,opt,name=code,proto3,enum=cbridgenode.ErrorCode" json:"code,omitempty"`
	Msg  string    `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *CbridgeNodeError) Reset() {
	*x = CbridgeNodeError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cbridge_node_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CbridgeNodeError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CbridgeNodeError) ProtoMessage() {}

func (x *CbridgeNodeError) ProtoReflect() protoreflect.Message {
	mi := &file_cbridge_node_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CbridgeNodeError.ProtoReflect.Descriptor instead.
func (*CbridgeNodeError) Descriptor() ([]byte, []int) {
	return file_cbridge_node_proto_rawDescGZIP(), []int{6}
}

func (x *CbridgeNodeError) GetCode() ErrorCode {
	if x != nil {
		return x.Code
	}
	return ErrorCode_INVALID
}

func (x *CbridgeNodeError) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_cbridge_node_proto protoreflect.FileDescriptor

var file_cbridge_node_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x6e, 0x6f, 0x64,
	0x65, 0x22, 0x9e, 0x01, 0x0a, 0x0d, 0x43, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x26, 0x0a, 0x0f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x6f, 0x64,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65,
	0x6c, 0x61, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x63, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0b, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x62, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x64, 0x62, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x22, 0xc8, 0x03, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x3b, 0x0a, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x62, 0x72,
	0x69, 0x64, 0x67, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x0b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x3b, 0x0a, 0x0c, 0x77, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x62, 0x72, 0x69, 0x64, 0x67,
	0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x0b, 0x77, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x24,
	0x0a, 0x0e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x5f, 0x67, 0x61, 0x73, 0x5f, 0x67, 0x77, 0x65, 0x69,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x47, 0x61, 0x73,
	0x47, 0x77, 0x65, 0x69, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x65, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x66, 0x65, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12,
	0x24, 0x0a, 0x0e, 0x67, 0x61, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x67, 0x61, 0x73, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x67, 0x61, 0x73, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x5f, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0f, 0x67, 0x61, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61,
	0x6c, 0x12, 0x4a, 0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63,
	0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x10, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x76, 0x0a,
	0x0b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x23, 0x0a, 0x0d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x44, 0x65,
	0x63, 0x69, 0x6d, 0x61, 0x6c, 0x22, 0xb1, 0x01, 0x0a, 0x0b, 0x57, 0x61, 0x74, 0x63, 0x68, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0f, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x65, 0x6c, 0x61,
	0x79, 0x12, 0x26, 0x0a, 0x0f, 0x6d, 0x61, 0x78, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x64,
	0x65, 0x6c, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x6d, 0x61, 0x78, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x12, 0x2e, 0x0a, 0x13, 0x66, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x22, 0x86, 0x01, 0x0a, 0x10, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1b,
	0x0a, 0x09, 0x67, 0x61, 0x73, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x67, 0x61, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x61,
	0x64, 0x64, 0x5f, 0x67, 0x61, 0x73, 0x5f, 0x67, 0x77, 0x65, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0a, 0x61, 0x64, 0x64, 0x47, 0x61, 0x73, 0x47, 0x77, 0x65, 0x69, 0x12, 0x33, 0x0a,
	0x16, 0x61, 0x64, 0x64, 0x5f, 0x67, 0x61, 0x73, 0x5f, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74,
	0x65, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x13, 0x61,
	0x64, 0x64, 0x47, 0x61, 0x73, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x52, 0x61, 0x74,
	0x69, 0x6f, 0x22, 0xe9, 0x02, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12,
	0x1f, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x73, 0x12, 0x33, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e,
	0x63, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x2d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x19, 0x2e, 0x63, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x41, 0x64, 0x64, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x66, 0x65, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08,
	0x66, 0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x6f, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x74, 0x6f, 0x41, 0x64, 0x64,
	0x72, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x22, 0x50,
	0x0a, 0x10, 0x43, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x2a, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x16, 0x2e, 0x63, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x2a, 0xa1, 0x02, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x19, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x01, 0x12, 0x1d,
	0x0a, 0x19, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x52, 0x4d, 0x45, 0x44, 0x10, 0x02, 0x12, 0x1c, 0x0a,
	0x18, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x52, 0x45, 0x46, 0x55, 0x4e, 0x44, 0x45, 0x44, 0x10, 0x03, 0x12, 0x25, 0x0a, 0x21, 0x54,
	0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x54,
	0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54,
	0x10, 0x04, 0x12, 0x27, 0x0a, 0x23, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x49,
	0x4e, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x12, 0x23, 0x0a, 0x1f, 0x54,
	0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43,
	0x4f, 0x4e, 0x46, 0x49, 0x52, 0x4d, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x06,
	0x12, 0x22, 0x0a, 0x1e, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x52, 0x45, 0x46, 0x55, 0x4e, 0x44, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49,
	0x4e, 0x47, 0x10, 0x07, 0x2a, 0x58, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x4f, 0x55, 0x54, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x52, 0x41, 0x4e,
	0x53, 0x46, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x10, 0x02, 0x2a, 0x18,
	0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x2d, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x63, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2d, 0x67, 0x6f, 0x2f, 0x63, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x6e, 0x6f, 0x64,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cbridge_node_proto_rawDescOnce sync.Once
	file_cbridge_node_proto_rawDescData = file_cbridge_node_proto_rawDesc
)

func file_cbridge_node_proto_rawDescGZIP() []byte {
	file_cbridge_node_proto_rawDescOnce.Do(func() {
		file_cbridge_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_cbridge_node_proto_rawDescData)
	})
	return file_cbridge_node_proto_rawDescData
}

var file_cbridge_node_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_cbridge_node_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_cbridge_node_proto_goTypes = []interface{}{
	(TransferStatus)(0),      // 0: cbridgenode.TransferStatus
	(TransferType)(0),        // 1: cbridgenode.TransferType
	(ErrorCode)(0),           // 2: cbridgenode.ErrorCode
	(*CBridgeConfig)(nil),    // 3: cbridgenode.CBridgeConfig
	(*ChainConfig)(nil),      // 4: cbridgenode.ChainConfig
	(*TokenConfig)(nil),      // 5: cbridgenode.TokenConfig
	(*WatchConfig)(nil),      // 6: cbridgenode.WatchConfig
	(*TransactorConfig)(nil), // 7: cbridgenode.TransactorConfig
	(*Transfer)(nil),         // 8: cbridgenode.Transfer
	(*CbridgeNodeError)(nil), // 9: cbridgenode.CbridgeNodeError
}
var file_cbridge_node_proto_depIdxs = []int32{
	4, // 0: cbridgenode.CBridgeConfig.chain_config:type_name -> cbridgenode.ChainConfig
	5, // 1: cbridgenode.ChainConfig.token_config:type_name -> cbridgenode.TokenConfig
	6, // 2: cbridgenode.ChainConfig.watch_config:type_name -> cbridgenode.WatchConfig
	7, // 3: cbridgenode.ChainConfig.transactor_config:type_name -> cbridgenode.TransactorConfig
	0, // 4: cbridgenode.Transfer.status:type_name -> cbridgenode.TransferStatus
	1, // 5: cbridgenode.Transfer.type:type_name -> cbridgenode.TransferType
	2, // 6: cbridgenode.CbridgeNodeError.code:type_name -> cbridgenode.ErrorCode
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_cbridge_node_proto_init() }
func file_cbridge_node_proto_init() {
	if File_cbridge_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cbridge_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CBridgeConfig); i {
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
		file_cbridge_node_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChainConfig); i {
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
		file_cbridge_node_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenConfig); i {
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
		file_cbridge_node_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchConfig); i {
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
		file_cbridge_node_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactorConfig); i {
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
		file_cbridge_node_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transfer); i {
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
		file_cbridge_node_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CbridgeNodeError); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cbridge_node_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cbridge_node_proto_goTypes,
		DependencyIndexes: file_cbridge_node_proto_depIdxs,
		EnumInfos:         file_cbridge_node_proto_enumTypes,
		MessageInfos:      file_cbridge_node_proto_msgTypes,
	}.Build()
	File_cbridge_node_proto = out.File
	file_cbridge_node_proto_rawDesc = nil
	file_cbridge_node_proto_goTypes = nil
	file_cbridge_node_proto_depIdxs = nil
}
