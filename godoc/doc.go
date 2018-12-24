/*
If you have a file named "doc.go" in your package directory, it becomes the Overview section when rendered by the godoc tool.

Notice how the first line is rendered as the summary in the package index.

  If you indent, you will get an block section in the documentation

    type Foo struct
	{
	   Bar string
	   Baz int
	}

Be sure that the "package" declaration is last in doc.go and that there are no blank lines between this comment section and
the "package godoc" declaration.

Failure to heed this advice may result in the Overview section being blank.

Starting the godoc Server

godoc -http=":port"

  Exmple:

  godoc -http=":6060"&

If is hanndy to throw the godoc sever in the background.  The godoc server will pick up modifications to your source code when you hit 
refresh in your browser.


Notice how we start new sections

By entering a line of text followed by a paragrahph of information that spans a few lines. The paragraph heading should not contain ending 
punctuation or else it will be mistaken by the parser as simple text and not a header.

Creating Examples

The "os" package has a number of examples, see https://golang.org/pkg/os/#pkg-examples.  You may have wondered how to do this yourself.
Similar to creating a regular unit test;

  func TestSomething(t *testing.T){
     // do stuff
  }

You can provide a test that will be rendered in the documentation by using "ExampleSomething".

  func ExampleSomething() {
    // do stuff

  }

You can even run the godoc server in "playground" mode.  In this mode, your examples can be modified and run by the user similiar to 
https://play.golang.org/.


For more information about godoc, see;

  * https://godoc.org/github.com/fluhus/godoc-tricks
  * https://blog.golang.org/godoc-documenting-go-code

*/
package godoc
