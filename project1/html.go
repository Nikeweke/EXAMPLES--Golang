
package main


var (

      t = `
      <html>
            <head>
            <title>HOME</title>
            </head>

      <body>
        <h2>HALLO</h2>
        <form method="post">
           <label>Email:</label>
           <input name="email" type="text">
           <br>
           <label>Password:</label>
           <input name="password" type="text">
           <button name="btn" value="enter" type="submit">OK</button>
        </form>
      </body>
      </html> `

     a = `<html>
          <head>
          <title>NOT HOME</title>
          </head>
          <body>
            <h2>%s, (%s)</h2>
          </body>
          </html> `
  )
