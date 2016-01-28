package db_test

import (
	. "github.com/bentrevor/calhoun/db"

	"testing"

	. "github.com/bentrevor/calhoun/spec-helper"
)

func TestPhotoFS_Filepath(t *testing.T) {
	// photos will be stored in /srv/images, in directories based on the md5 hash of their id
	// padded in front with 0's to 12 decimal places

	// $ echo -n "000000000012" | md5sum
	// 9ed63b492437de85736cb562f91f203c  -
	want := "/fake/srv/path/9e/d6/3b492437de85736cb562f91f203c"

	Describe("PhotoFS: filepath")
	photo := Photo{Id: 12}
	fs := NewPhotoFS("/fake/srv/path")

	It("takes the md5 hash of (photo_id padded in front with 0s to 12 places)")
	AssertEquals(t, want, fs.PhotoFilepath(photo))
}
