# Marp template

- Template for SLFX Presentations.
- Include easily include code snippets from actual files.
- Build & presentation step prepared.
- Rose pine themes included.

https://github.com/schuettflix/marp-template/assets/14948823/c6b639fc-1aa7-4ec0-9ba0-318a81ce1478

## Usage

```bash
# Prepare
git clone git@github.com:schuettflix/marp-template.git
cd marp-template
bun i

# Write your presentation in realtime
bun run serve
bun run watch

# When you are done, present with
bun run serve
# Or
bun run open
```

Other useful commands

```bash
# Instead of html, you can build for pdf or images
bun run build:pdf
bun run build:images
```

## Customise

Most stuff (footer, header, etc.) can be tweaked in the header of the `presentation.md` file.

### Theme

There are 3 built in themes.

```yal
theme: rose-pine
theme: rose-pine-moon
theme: rose-pine-dawn
```
