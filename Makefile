.PHONY: new
new:
	@if [ -z "$(YEAR)" ] || [ -z "$(DAY)" ]; then \
		echo "Usage: make new YEAR=<year> DAY=<day>"; \
		echo "Example: make new YEAR=2021 DAY=2"; \
		exit 1; \
	fi; \
	mkdir -p $(YEAR)/day$(shell printf "%02d" $(DAY)); \
	cp template/day.go $(YEAR)/day$(shell printf "%02d" $(DAY))/day.go; \
	sed -e 's/package dayXX/package day$(shell printf "%02d" $(DAY))/' \
	    -e 's/const year = 1970/const year = $(YEAR)/' \
	    -e 's/const day = 0/const day = $(DAY)/' \
	    $(YEAR)/day$(shell printf "%02d" $(DAY))/day.go > $(YEAR)/day$(shell printf "%02d" $(DAY))/day.go.tmp && \
	mv $(YEAR)/day$(shell printf "%02d" $(DAY))/day.go.tmp $(YEAR)/day$(shell printf "%02d" $(DAY))/day.go; \
	awk '/github.com\/tofoss\/aoc-go\/pkg\/registry/ {print "\t_ \"github.com/tofoss/aoc-go/$(YEAR)/day$(shell printf "%02d" $(DAY))\""} {print}' main.go > main.go.tmp && mv main.go.tmp main.go; \
	echo "Created $(YEAR)/day$(shell printf "%02d" $(DAY))/day.go and added import to main.go"
