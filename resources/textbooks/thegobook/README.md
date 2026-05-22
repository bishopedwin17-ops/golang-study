# Introducing Go — Textbook Integration

This folder contains integration metadata for the textbook PDF imported into this repository.

## Asset location
- `resources/textbooks/thegobook.pdf` — the downloaded textbook file.

## Purpose
- Keep the textbook available as a reference asset.
- Map the book's chapters to the existing Go study repository topics.
- Provide a lightweight integration that does not affect Go package compilation.

## How to use
- Open `resources/textbooks/thegobook.pdf` in any PDF viewer.
- Use `toc.md` for a quick chapter reference.
- Use `mapping.md` to connect textbook chapters with repo folders and exercises.

## Notes
- This integration is intentionally asset-first: the PDF is stored separately from source code.
- The repository remains compilable because these files are Markdown/PDF assets only.
- If you want, I can next add chapter summaries under this folder to make the book content searchable.
