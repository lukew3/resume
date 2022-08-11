require 'json'
require 'erb'
require 'ostruct'

file = File.read('./resume.json')
data_hash = JSON.parse(file)
latex_render = ERB.new(File.read('./template.tex.erb')).result(OpenStruct.new(data_hash).instance_eval { binding })
File.open('./resume.tex', 'w') { |f| f.write(latex_render) }
