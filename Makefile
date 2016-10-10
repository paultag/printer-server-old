CSS_FILES = \
			src/stylesheets/all.css \
			lib/weather-icons/sass/weather-icons.css \
			stylesheet.css

SCSS = scss
OUTPUT = output
OUTPUT_FONTS = $(OUTPUT)/font
OUTPUT_CSS = $(OUTPUT)/stylesheets

.SUFFIXES:
.SUFFIXES: .scss .css

all: clean build

build: $(CSS_FILES)

clean:
	rm -fv $(CSS_FILES)
	rm -fvr $(OUTPUT)

.scss.css:
	$(SCSS) $< $@

output: build
	mkdir -p $(OUTPUT) $(OUTPUT_FONTS) $(OUTPUT_CSS)
	cp -v src/font/* $(OUTPUT_FONTS)
	cp -v lib/weather-icons/font/* $(OUTPUT_FONTS)
	cp -v $(CSS_FILES) $(OUTPUT_CSS)

.PHONY: output
