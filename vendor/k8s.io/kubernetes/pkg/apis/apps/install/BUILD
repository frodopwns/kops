package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "go_default_library",
    srcs = ["install.go"],
    tags = ["automanaged"],
    deps = [
        "//pkg/apimachinery/announced:go_default_library",
        "//pkg/apis/apps:go_default_library",
        "//pkg/apis/apps/v1beta1:go_default_library",
    ],
)
