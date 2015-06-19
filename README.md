# svg2pdf
Creates layered PDF file from series of SVG files

## About

This app is written in GO uses modified library [gofpdf](https://github.com/kpawlik/gofpdf).
svg2pdf converts series of SVG files into layered PDF file, each SVG file is put in separate layer.
Name of file become name of the layer.

This app was created to convert only SVG files creaed by one specific application, so library doesn't support all SVG syles.


Usage:
```
svg2pdf.exe [OPTIONS] LIST_OF_SVG_FILES
```

Options:

- `format` - PDF format name. Default `A4`. Supported formats `A3, A4, A5`
- `orientation` - PDF page orientation. Default `P`. Supported values `P`, `L`. (`P` - portrail, `L` - landscape)
- `out` - PDF output file path. Default  `result.pdf`.
- `dir` - Path to directory which contains SVG files. All SVG files from this library will be added to PDF. This parameter will be ignored if list of svg files is present.
- `linew` - float. Base line width.
- `scale` - float. Base scale factor.
- `time` - boolean. Prints some time statistics.
