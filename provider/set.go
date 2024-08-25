package provider

func tagSet(tags []string) []string {
	tagMap := map[string]struct{}{}

	for _, tag := range tags {
		tagMap[tag] = struct{}{}
	}

	set := []string{}

	for tag := range tagMap {
		set = append(set, tag)
	}

	return set
}
