#!/bin/bash

# 功能：生成proto文件
#
# 由于bcloud不支持生成go版本的proto，且集成编译时也无法直接拉取外网github.com/golang/protobuf包
# 所以直接在本地生成proto文件，上传代码库；
# 需要注意本地protoc-gen-go的版本需要与mod中的版本保持一致
#
# require (
#  	github.com/golang/protobuf v1.4.0
# )
#
# cd github.com/golang/protobuf/protoc-gen-go
# git tag -l
# git checkout v1.4.0
# go build
# go install
#

proto_dir=github.com/yincongcyincong/BasicStudy/proto/api

# 需要回退到根目录执行
function gen() (
    cd ../../../
    pwd
    for proto in $(find "$proto_dir" -name '*.proto'); do
        echo "$proto"
        protoc --go_out=. "$proto"
    done

    # 扩展类型，array类型字段
)

function gen_one() {
    proto="$proto_dir/$2"
    cd ../../../
    pwd
    echo "$proto.proto"
    protoc --go_out=. "$proto.proto"

    go_file="$proto.pb.go"
    if [ ! -f "$go_file" ];then
         # 截取方法名
        method=${proto##*/}
        go_file=$(find "$proto_dir" -name "$method.pb.go")
    fi
    echo "go_file: $go_file"
    # 生成 mcpack 标签
}

function remove_omitempty() {
    proto=$1
    attr_name=$2
    sed -i '.bak' 's/json:"\([^,]*\),omitempty\"/json:"\1"/g' "$proto"
}

function replace_attr() {
    proto=$1
    attr_name=$2
    echo "replace attr:[$attr_name] in file:[$proto]"
    sed -i'.bak' "s/\(^	$attr_name *\*\)string\(.*\)/\1interface{}\2/g" "$proto"
    sed -i'.bak' "s/\( Get$attr_name() \)string {/\1interface{} {/g" "$proto"
}

function clean() {
    find ./bin/proto.sh -name '*.pb.go' | xargs rm
}

function clean_bak() {
    find ./bin/proto.sh -name '*.pb.go.bak' | xargs rm
}

function usage() {
    echo "usage: $(basename "$0") [gen|gen_one|clean|clean_bak|replace|replace_required]"
}

function main() {
    action=$1
    case "X${action}" in
        Xgen)
            gen
            ;;
        Xgen_one)
            gen_one "$@"
            ;;
        Xclean)
            clean
            ;;
        Xclean_bak)
            clean_bak
            ;;
        Xreplace)
            shift
            replace "$@"
            ;;
        Xreplace_required)
            shift
            replace_required "$@"
            ;;
        *)
            usage
            ;;
    esac
}

main "$@"
