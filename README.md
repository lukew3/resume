# lukew3/resume
Tracking how my resume changes using JSONresume. Feel free to fork or copy my github workflows and watch github actions automatically validate and build your resume on release.

The latest pdf build can be downloaded [here](https://github.com/lukew3/resume/releases/latest/download/lukew3_resume.pdf)

### How to use it for your own resume
1. Fork the repo
2. Click on the actions tab of the forked repo and enable actions
3. Make desired changes to the `resume.json` file and/or `template.tex.erb`
5. Click on the Releases sidebar on your Github fork, draft a new release, add a tag like `0.0.1` and click "Publish release". Your generated pdf will appear in the assets section of the release in around 2.5 minutes.

Make a release every time that you want to publish a new version of your resume. You can then download the pdf from the assets section or have a permanent link to the latest release with a url of the format `https://github.com/{yourUsername}/resume/releases/latest/download/{yourUsername}_resume.pdf` where `{yourUsername}` is replaced with your Github username.

### How it works
The content of my resume is stored in the [`resume.json`](https://github.com/lukew3/resume/blob/main/resume.json) file, following the schema defined by [jsonresume.org](https://jsonresume.org/schema/). When a release is created, the [`release.yml`](https://github.com/lukew3/resume/blob/main/.github/workflows/release.yml) workflow is ran.

#### Release.yml
The following are the steps taken by [`release.yml`](https://github.com/lukew3/resume/blob/main/.github/workflows/release.yml):
1. Validates json resume using [lukew3/validate-json-resume-action](https://github.com/lukew3/validate-json-resume-action).
2. Fill [`template.tex.erb`](https://github.com/lukew3/resume/blob/main/template.tex.erb) with contents of [`resume.json`](https://github.com/lukew3/resume/blob/main/resume.json) using [lukew3/json-fill-erb-action](https://github.com/lukew3/json-fill-erb-action).
3. `resume.tex` which was created by step 2 is built to `resume.pdf` with [xu-cheng/latex-action](https://github.com/xu-cheng/latex-action).
4. The resulting pdf is uploaded to the latest release with the file name `{username}_resume.pdf`.

### Building on Your PC
The release workflow will automatically build a pdf on release. If you want to render locally, you can use the `build.rb` file [here](https://github.com/lukew3/resume/blob/local-build/build.rb) and then run `pdflatex resume.tex`.
