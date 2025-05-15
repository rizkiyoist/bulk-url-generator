This is a simple Go program to generate URLs based on a given template.

The program takes these parameters:

* `baseUrl`: the base URL to use for the generated URLs
* `from`: the starting number for the generated URLs
* `to`: the ending number for the generated URLs
* `isPadded`: a boolean indicating whether to use padding for the numbers (e.g. 001, 002)
* `padDigit`: the padding digit to use (e.g. 4 means 0001, 0021)

The program generates the URLs and writes them to a file with the current timestamp as the filename.

To use this program, simply run it with the desired parameters: