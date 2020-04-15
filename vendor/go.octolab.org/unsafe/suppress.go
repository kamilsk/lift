package unsafe

// DoSilent accepts a result of
// * fmt.Fprint* function family
// * io.Copy* and io.Read* function family
// * io.Writer interface
// and allows to ignore it.
//
//  unsafe.DoSilent(fmt.Fprintln(writer, "ignore the result"))
//
func DoSilent(interface{}, error) {}

// Ignore accepts an error and allows to ignore it.
//
//  unsafe.Ignore(
//  	template.Must(template.New("html").Parse(content)).Execute(writer, data),
//  )
//
func Ignore(error) {}
