package day5

import (
	"slices"
	"strconv"
	"strings"
)

func Solve2(input string) (any, error) {
	type idRange struct {
		from, to int
	}

	var (
		idRangesRaw = strings.Split(strings.Split(input, "\n\n")[0], "\n")
		idRanges    = make([]idRange, 0, len(idRangesRaw))
	)

	for _, idRangeRaw := range idRangesRaw {
		bounds := strings.Split(idRangeRaw, "-")
		from, err := strconv.Atoi(bounds[0])
		if err != nil {
			return nil, err
		}

		to, err := strconv.Atoi(bounds[1])
		if err != nil {
			return nil, err
		}

		idRanges = append(idRanges, idRange{
			from: from,
			to:   to,
		})
	}

	for {
		var (
			fst, snd int
			found    bool
		)

		for f, range1 := range idRanges {
			for s, range2 := range idRanges {
				if f >= s {
					continue
				}

				if range1.from >= range2.from && range1.from <= range2.to ||
					range1.to >= range2.from && range1.to <= range2.to {
					fst = f
					snd = s
					found = true
				}
			}
		}

		if !found {
			break
		}

		newRange := idRange{
			from: min(idRanges[fst].from, idRanges[snd].from),
			to:   max(idRanges[fst].to, idRanges[snd].to),
		}

		idRanges = slices.Delete(idRanges, max(fst, snd), max(fst, snd)+1)
		idRanges = slices.Delete(idRanges, min(fst, snd), min(fst, snd)+1)
		idRanges = append(idRanges, newRange)
	}

	var ans int

	for _, r := range idRanges {
		ans += r.to - r.from + 1
	}

	return ans, nil
}
