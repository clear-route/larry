package parse

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// struct for the git tags
type GitTags struct {
	rawTags []byte
}

type VersionTag struct {
	Major int
	Minor int
	Patch int
}

func (v *VersionTag) String() string {
	return fmt.Sprintf("v%d.%d.%d", v.Major, v.Minor, v.Patch)
}

func (v *VersionTag) IncrementMajor() string {
	fmt.Println("old major:", v.Major)
	v.Major = v.Major+1
	fmt.Println("new major:", v.Major)
	return v.String()
}

func (v *VersionTag) IncrementMinor() string {
	v.Minor++
	return v.String()
}

func (v *VersionTag) IncrementPatch() string {
	v.Patch++
	return v.String()
}

func Parse(tag string) VersionTag {
	res := VersionTag{}
	re := regexp.MustCompile(`^v\d+\.\d+\.\d+$`)
	if re.MatchString(tag) {
		segments := strings.Split(strings.TrimPrefix(tag, "v"), ".")
		var err error
		res.Major, err = strconv.Atoi(segments[0])
		if err != nil {
			panic(err)
		}
		res.Minor, err = strconv.Atoi(segments[1])
		if err != nil {
			panic(err)
		}
		res.Patch, err = strconv.Atoi(segments[2])
		if err != nil {
			panic(err)
		}
	}
	return res
}

// gets the tags from the git repository
func (g *GitTags) getTags() {
	cmd := exec.Command("git", "tag")
	bTags, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	g.rawTags = bTags
}

// returns the last tag
func (g *GitTags) LastTag() VersionTag {
	g.getTags()
	rawTagsStr := string(g.rawTags)
	tags := strings.Split(string(rawTagsStr), "\n")
	lastTag := tags[len(tags)-2]
	t := Parse(lastTag)
	return t
}
