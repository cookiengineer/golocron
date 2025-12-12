package golocron

import "embed"

//go:embed public/*
var Filesystem embed.FS
