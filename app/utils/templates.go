package utils

func GetPage(content string) string {
	resultPage := `<!DOCTYPE html>
		<html lang="en">
		<head>
		   <meta charset="UTF-8"/>
		   <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
		   <meta http-equiv="X-UA-Compatible" content="IE=7"/>
		   <link rel="stylesheet" href="https://use.fontawesome.com/releases/v5.6.3/css/all.css"
				 integrity="sha384-UHRtZLI+pbxtHCWp1t77Bi1L4ZtiqrqD80Kn4Z8NTSRyMA2Fd33n5dQ8lWUE00s/" crossorigin="anonymous"/>
		   <link
				   href="https://fonts.googleapis.com/css?family=Nunito:200,200i,300,300i,400,400i,600,600i,700,700i,800,800i,900,900i"
				   rel="stylesheet">
		   <link rel="stylesheet" href="https://bootswatch.com/4/journal/bootstrap.min.css"/>
		   <title>Currency converter</title>
		</head>
		
		<body>
		<div class="row mt-5">
		   <div class="col-md-6 m-auto">
			   <h2>Currency Converter</h2>
			   <div class="card card-body border-success">
				   <p class="lead">Convert the currency</p>
				   <form class="form-inline" action="/convert" method="post">
					   <div class="form-group mb-2">
						   <input type="number" class="form-control" name="amount" placeholder="0.00" required/>
					   </div>
					   <div class="form-group mx-sm-3 mb-2">
						   <select class="form-control" name="from_currency" required>
							   <option>NGN</option>
							   <option>GHS</option>
							   <option>KSH</option>
						   </select>
					   </div>
					   <div class="form-group mx-sm-3 mb-2">
						   <label>convert to &emsp;</label>
						   <select class="form-control" name="to_currency" required>
							   <option>NGN</option>
							   <option>GHS</option>
							   <option>KSH</option>
						   </select>
					   </div>
					   <button type="submit" class="btn btn-success mb-2">Convert</button>
				   </form>` + content + `
			   </div>
		   </div>
		</div>
		</div>
		</body>
		
		</html>
		`
	return resultPage
}
