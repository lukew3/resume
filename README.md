# lukew3/resume
Tracking how my latex resume changes and building a hosted pdf with Github Actions. Feel free to fork or copy my github workflows and watch github actions automatically validate and build your resume on release.

Inspired by the [resumake.io](https://github.com/saadq/resumake.io) website and created out of the desire to make an automatic latex resume build system utilizing git version control.

The latest pdf build can be downloaded [here](https://github.com/lukew3/resume/releases/latest/download/lukew3_resume.pdf)

### How to use it for your own resume
1. Fork the repo
2. Click on the actions tab of the forked repo and enable actions
3. Click on the settings tab, go to actions->general on the sidebar, scroll down to workflow permissions, set to "Read and write permissions", and save.
4. Make desired changes to `resume.tex`
5. Click on the Releases sidebar on your Github fork, draft a new release, add a tag like `0.0.1` and click "Publish release". Your generated pdf will appear in the assets section of the release in around 2.5 minutes.

Make a release every time that you want to publish a new version of your resume. You can then download the pdf from the assets section or have a permanent link to the latest release with a url of the format `https://github.com/{yourUsername}/resume/releases/latest/download/{yourUsername}_resume.pdf` where `{yourUsername}` is replaced with your Github username.

### How it works
The content of the resume is stored in the [`resume.tex`](https://github.com/lukew3/resume/blob/main/resume.tex) file. When a release is created, the [`release.yml`](https://github.com/lukew3/resume/blob/main/.github/workflows/release.yml) workflow is ran.

#### Release.yml
The following are the steps taken by [`release.yml`](https://github.com/lukew3/resume/blob/main/.github/workflows/release.yml):
1. `resume.tex` is built to `resume.pdf` with [xu-cheng/latex-action](https://github.com/xu-cheng/latex-action).
2. The resulting pdf is uploaded to the latest release with the file name `{username}_resume.pdf`.

### Building on Your PC
The release workflow will automatically build a pdf on release. If you want to render locally, you can use the `build.rb` file [here](https://github.com/lukew3/resume/blob/local-build/build.rb) and then run `pdflatex resume.tex`.
