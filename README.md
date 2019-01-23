# go-lambda-email

AWS Lambda function written in Go to deliver emails that are submitted via 
a POST request from a form.

The following are combined to send the emails:
- ```Go``` program
- AWS API code from GitHub
- AWS Lambda behind an API Gateway
- AWS Simple Email Service (SES)

## Building

### In Windows:
```set GOOS=linux```

```go build -o main cmd\go-lambda-email\main.go```

```%GOROOT%\bin\build-lambda-zip.exe -o main.zip main```

## Configure AWS
The following will be done from the [AWS Console](https://console.aws.amazon.com/console/home).

The examples are running in the ```US East (N. Virginia)``` aka ```us-east-1``` data center.

<img src="https://www.dropbox.com/s/g5at0724x3wwhw3/aws-console-overview-annotated.png?raw=true" alt="aws" width="700"/>

### Configure Simple Email Service (SES)
1. From the AWS console, open **_Simple Email Service_** under **_Customer Engagement_**.
2. Open **_Email Addresses_** under **_Identity Management_**.
  <img src="https://www.dropbox.com/s/3hrx2hg4hmwukq4/aws-ses-overview.png?raw=true" alt="ses" width="700"/>
3. Click on **_Verify a New Email Address_** and follow the steps to add a new entry.

### Create the Lambda function
1. From the AWS console, open open **_Lambda_** under **_Compute_**
2. Click on **_Create Function_** and use the following settings:<br>
  **_Create Function_** -> **_Author from scratch_**
  1. **```Name```**: anything you want - this example uses `go-lambda-email`
  2. **```Runtime```**: Go 1.x
  3. **```Role```**: Create a custom role
    - A new window/tab should open - follow these [steps](#iam-custom-role-for-ses) to create an IAM role
  4. Click **_Create function_** which should forward you to summary for the new function
3. Enter keys/values for the required ```ENV``` variables
  - **```AWS_DATACENTER```**: ```us-east-1```
  - **```TO_EMAIL```**: ```name@example.com```  
  <img src="https://www.dropbox.com/s/7gjzyb7u4ekbmd7/aws-lambda-env.png?raw=true" alt="ses" width="600"/>
  
  **Note:** The **```TO_EMAIL```** value should be set to a validated SES address
4. Under **_Function code_**, upload the previously created .zip
5. Change the **_Handler_** to ```main```

  <img src="https://www.dropbox.com/s/o5ioyjc8xykcwxe/aws-lambda-function.png?raw=true" alt="ses" width="600"/>  

6. Click **_Save_**
7. Configure an [API Gateway](#api-gateway)

#### IAM custom role for SES

This role will grant the Lambda access to SES.

From the new IAM role screen, enter the following:

1. **```IAM Role```**: Create a new IAM Role
2. **```Role Name```**: anything you want - this example uses `lambda_basic_execution_with_ses`
3. Expand the **View Policy Document** section and click **_edit_**
4. Add the following bold section to the existing "Statement" (don't forget the comma)
<pre><code>
  {
    "Version": "2012-10-17",
    "Statement": [
      {
        "Effect": "Allow",
        "Action": [
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:PutLogEvents"
        ],
        "Resource": "arn:aws:logs:*:*:*"
      }<b><i>,
      {
        "Effect": "Allow",
        "Action": "ses:*",
        "Resource": "*"
      }</i></b>
    ]
  }
</code></pre>
5. Click **_Allow_** to finish.
  [Back](#create-the-lambda-function) to creating the Lambda.
  
  
#### API Gateway

From the Lambda summary screen, click on **_API Gateway_** under **_Designer_**.

Enter the following:

1. **```API```**: ```Create a new API```
2. **```Security```**: ```Open with API key```
3. Expand **_Additional settings_**
4. Name the API (optional - a default name should be filled in)
5. Name the deployment stage (optional - a default name should be filled in)
6. Back on the Lambda summary screen, click on **_Save_**
  - An **API endpoint** and **API key** should be generated
  
  <img src="https://www.dropbox.com/s/2mpg0ttxi7rypxt/aws-api-gateway.png?raw=true" alt="ses" width="600"/>

## Testing the function

Create a POST request (screenshots use [Postman](https://www.getpostman.com/)).

Be sure to add a ```x-api-key``` header with the value of the generated key from your API Gateway.

### Client Responses

Success (```200``` response):

<img src="https://www.dropbox.com/s/5dtuy00pmyhykq7/postman-testing-success.png?raw=true" alt="ses" width="600"/>

Failure (```400``` response and returned error message):

<img src="https://www.dropbox.com/s/t8hem9ndqew660m/postman-testing-error.png?raw=true" alt="ses" width="600"/>

### Delivered Email

The following is an example of the delivered email:

<img src="https://www.dropbox.com/s/skgph207o9x3xgg/success-inbox.png?raw=true" alt="ses" width="600"/>
