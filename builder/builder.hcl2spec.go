// Code generated by "packer-sdc mapstructure-to-hcl2"; DO NOT EDIT.

package builder

import (
	"github.com/davejbax/packer-plugin-arm/config"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	FileChecksum              *string                  `mapstructure:"file_checksum" required:"true" cty:"file_checksum" hcl:"file_checksum"`
	FileChecksumURL           *string                  `mapstructure:"file_checksum_url" cty:"file_checksum_url" hcl:"file_checksum_url"`
	FileChecksumType          *string                  `mapstructure:"file_checksum_type" cty:"file_checksum_type" hcl:"file_checksum_type"`
	FileUrls                  []string                 `mapstructure:"file_urls" cty:"file_urls" hcl:"file_urls"`
	FileUnarchiveCmd          []string                 `mapstructure:"file_unarchive_cmd" cty:"file_unarchive_cmd" hcl:"file_unarchive_cmd"`
	TargetPath                *string                  `mapstructure:"file_target_path" cty:"file_target_path" hcl:"file_target_path"`
	TargetExtension           *string                  `mapstructure:"file_target_extension" cty:"file_target_extension" hcl:"file_target_extension"`
	TmpDirLocation            *string                  `mapstructure:"file_tmp_dir_location" cty:"file_tmp_dir_location" hcl:"file_tmp_dir_location"`
	ImagePath                 *string                  `mapstructure:"image_path" required:"true" cty:"image_path" hcl:"image_path"`
	ImageSize                 *string                  `mapstructure:"image_size" cty:"image_size" hcl:"image_size"`
	ImageType                 *string                  `mapstructure:"image_type" cty:"image_type" hcl:"image_type"`
	ImageMountPath            *string                  `mapstructure:"image_mount_path" cty:"image_mount_path" hcl:"image_mount_path"`
	ImageBuildMethod          *string                  `mapstructure:"image_build_method" cty:"image_build_method" hcl:"image_build_method"`
	ImageSizeBytes            *uint64                  `mapstructure:"image_size_bytes" cty:"image_size_bytes" hcl:"image_size_bytes"`
	ImagePartitions           []config.FlatPartition   `mapstructure:"image_partitions" cty:"image_partitions" hcl:"image_partitions"`
	ImageChrootMounts         []config.FlatChrootMount `mapstructure:"image_chroot_mounts" cty:"image_chroot_mounts" hcl:"image_chroot_mounts"`
	AdditionalChrootMounts    []config.FlatChrootMount `mapstructure:"additional_chroot_mounts" cty:"additional_chroot_mounts" hcl:"additional_chroot_mounts"`
	ImageSetupExtra           [][]string               `mapstructure:"image_setup_extra" cty:"image_setup_extra" hcl:"image_setup_extra"`
	ImageChrootEnv            []string                 `mapstructure:"image_chroot_env" cty:"image_chroot_env" hcl:"image_chroot_env"`
	QemuBinarySourcePath      *string                  `mapstructure:"qemu_binary_source_path" required:"true" cty:"qemu_binary_source_path" hcl:"qemu_binary_source_path"`
	QemuBinaryDestinationPath *string                  `mapstructure:"qemu_binary_destination_path" cty:"qemu_binary_destination_path" hcl:"qemu_binary_destination_path"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"file_checksum":                &hcldec.AttrSpec{Name: "file_checksum", Type: cty.String, Required: false},
		"file_checksum_url":            &hcldec.AttrSpec{Name: "file_checksum_url", Type: cty.String, Required: false},
		"file_checksum_type":           &hcldec.AttrSpec{Name: "file_checksum_type", Type: cty.String, Required: false},
		"file_urls":                    &hcldec.AttrSpec{Name: "file_urls", Type: cty.List(cty.String), Required: false},
		"file_unarchive_cmd":           &hcldec.AttrSpec{Name: "file_unarchive_cmd", Type: cty.List(cty.String), Required: false},
		"file_target_path":             &hcldec.AttrSpec{Name: "file_target_path", Type: cty.String, Required: false},
		"file_target_extension":        &hcldec.AttrSpec{Name: "file_target_extension", Type: cty.String, Required: false},
		"file_tmp_dir_location":        &hcldec.AttrSpec{Name: "file_tmp_dir_location", Type: cty.String, Required: false},
		"image_path":                   &hcldec.AttrSpec{Name: "image_path", Type: cty.String, Required: false},
		"image_size":                   &hcldec.AttrSpec{Name: "image_size", Type: cty.String, Required: false},
		"image_type":                   &hcldec.AttrSpec{Name: "image_type", Type: cty.String, Required: false},
		"image_mount_path":             &hcldec.AttrSpec{Name: "image_mount_path", Type: cty.String, Required: false},
		"image_build_method":           &hcldec.AttrSpec{Name: "image_build_method", Type: cty.String, Required: false},
		"image_size_bytes":             &hcldec.AttrSpec{Name: "image_size_bytes", Type: cty.Number, Required: false},
		"image_partitions":             &hcldec.BlockListSpec{TypeName: "image_partitions", Nested: hcldec.ObjectSpec((*config.FlatPartition)(nil).HCL2Spec())},
		"image_chroot_mounts":          &hcldec.BlockListSpec{TypeName: "image_chroot_mounts", Nested: hcldec.ObjectSpec((*config.FlatChrootMount)(nil).HCL2Spec())},
		"additional_chroot_mounts":     &hcldec.BlockListSpec{TypeName: "additional_chroot_mounts", Nested: hcldec.ObjectSpec((*config.FlatChrootMount)(nil).HCL2Spec())},
		"image_setup_extra":            &hcldec.AttrSpec{Name: "image_setup_extra", Type: cty.List(cty.List(cty.String)), Required: false},
		"image_chroot_env":             &hcldec.AttrSpec{Name: "image_chroot_env", Type: cty.List(cty.String), Required: false},
		"qemu_binary_source_path":      &hcldec.AttrSpec{Name: "qemu_binary_source_path", Type: cty.String, Required: false},
		"qemu_binary_destination_path": &hcldec.AttrSpec{Name: "qemu_binary_destination_path", Type: cty.String, Required: false},
	}
	return s
}
