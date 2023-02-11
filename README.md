# lukew3/resume
Tracking how my resume changes using JSONresume. Feel free to fork and watch github actions automatically validate and build your resume on release.

The latest pdf build can be downloaded [here](https://github.com/lukew3/resume/releases/latest/download/lukew3_resume.pdf)


## How it works
The content of my resume is stored in the [`resume.json`](https://github.com/lukew3/resume/blob/main/resume.json) file, following the defined by [jsonresume.org](https://jsonresume.org/schema/). When a release is created, the [`release.yml`](https://github.com/lukew3/resume/blob/main/.github/workflows/release.yml) workflow is ran.

### Release.yml
The following are the steps taken by release.yml.
1. Validates json resume using [lukew3/validate-json-resume-action](https://github.com/lukew3/validate-json-resume-action).
2. Fill [`template.tex.erb`](https://github.com/lukew3/resume/blob/main/template.tex.erb) with contents of [`resume.json`](https://github.com/lukew3/resume/blob/main/resume.json) using [lukew3/json-fill-erb-action](https://github.com/lukew3/json-fill-erb-action).
3. `resume.tex` which was created by step 2 is built to `resume.pdf` with [xu-cheng/latex-action](https://github.com/xu-cheng/latex-action).
4. The resulting pdf is uploaded to the latest release with the file name `{username}_resume.pdf`.

## Building on Your PC
The release workflow will automatically build a pdf on release. If you want to render locally, you can use the `build.rb` file [here](https://github.com/lukew3/resume/blob/local-build/build.rb) and then run `pdflatex resume.tex`.
