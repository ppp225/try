// Package try provides retry functionality.
//     var value string
//     err := try.Do(func(attempt int) (error) {
//       var err error
//       value, err = SomeFunction()
//       return err
//     })
//     if err != nil {
//       log.Fatalln("error:", err)
//     }
package try
