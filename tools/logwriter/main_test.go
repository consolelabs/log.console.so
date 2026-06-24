package main

import "testing"

func TestSanitizeBody(t *testing.T) {
	cases := []struct {
		name, in, want string
	}{
		{
			name: "fence + front matter + trailing meta note (claude backend artifact)",
			in:   "```markdown\n---\ntitle: x\n---\n\n## TL;DR\nhi there\n```\n\nskipped: a shout-out section",
			want: "## TL;DR\nhi there",
		},
		{
			name: "plain body unchanged (api backend, clean)",
			in:   "## TL;DR\nclean body",
			want: "## TL;DR\nclean body",
		},
		{
			name: "leading front matter only",
			in:   "---\na: b\n---\n## Shipped\nthe thing",
			want: "## Shipped\nthe thing",
		},
		{
			name: "leading + trailing whitespace",
			in:   "\n\n## Next\nsoon\n\n",
			want: "## Next\nsoon",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := sanitizeBody(c.in); got != c.want {
				t.Errorf("sanitizeBody()\n got: %q\nwant: %q", got, c.want)
			}
		})
	}
}

func TestSignalGrouping(t *testing.T) {
	prs := []pr{
		{Title: "feat: a"}, {Title: "fix: b"},
	}
	prs[0].Repository.NameWithOwner = "consolelabs/mochi-api"
	prs[1].Repository.NameWithOwner = "consolelabs/mochi-api"
	g := groupByRepo(prs)
	if len(g) != 1 || len(g["consolelabs/mochi-api"]) != 2 {
		t.Fatalf("groupByRepo: got %v", g)
	}
}
