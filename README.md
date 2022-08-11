# resume
Tracking how my resume changes using JSONresume.

The latest pdf build can be downloaded [here](https://github.com/lukew3/resume/releases/latest/download/resume.pdf)

Resume built using https://resumake.io ([resumake repo](https://github.com/saadq/resumake.io))

## Building
First, install ruby and pdflatex. Then run `bundle install` to install the ruby dependencies.

To build, run 
```
ruby build.rb
```
and then run
```
pdflatex resume.tex
```