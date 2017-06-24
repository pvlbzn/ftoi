# ftoi: File to Image

A little tool which can render any file as an image. Particularly useful for reverse engineering.


## Examples

Two ([one](https://github.com/pvlbzn/ftoi/blob/master/samples/img.jpg), [two](https://github.com/pvlbzn/ftoi/blob/master/samples/img.png)) the same images by content, but different by compression method: jpg and png accordingly.

```
$ ftoi samples/img.jpg
...
$ ftoi samples/img.png
...
```

Will produce `out.png` file for each call.

Result for [jpg](https://github.com/pvlbzn/ftoi/blob/master/samples/jpg.png) and [png](https://github.com/pvlbzn/ftoi/blob/master/samples/png.png). So, despite that the content is visually the same, the actual bytes are not the same at all.

For example we can zoom into headers:

**JPG:**

![JPG header](https://github.com/pvlbzn/ftoi/blob/master/samples/jpeg_header.png)

**PNG:**

![PNG header](https://github.com/pvlbzn/ftoi/blob/master/samples/png_header.png)


## Install

```
$ go install https://github.com/pvlbzn/ftoi
```
