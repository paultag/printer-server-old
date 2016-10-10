CSS_FILES = src/stylesheets/all.css stylesheet.css

SCSS = scss
STATIC = static

.SUFFIXES:
.SUFFIXES: .scss .css

all: clean build

build: $(CSS_FILES)

clean:
	rm -fv $(CSS_FILES)

.scss.css:
	$(SCSS) $< $@

.PHONY: build
