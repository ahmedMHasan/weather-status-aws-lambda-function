# AWS Lambda Function for Dummy Weather Status and Current Time

This project demonstrates how to create an AWS Lambda function using Go to retrieve dummy weather information and current time. The Lambda function is triggered by an API Gateway endpoint and returns the weather details along with the current timestamp.

# Functionality

The Lambda function performs the following tasks:

- Receives a GET request from an API Gateway endpoint.
- Retrieves dummy weather information (in this example, it is hardcoded as "Sunny").
- Retrieves the current timestamp.
- Constructs a JSON response containing the weather information and timestamp.
- Returns the response to the API Gateway, which is then sent back to the client.

# Repository Structure

The repository has the following structure:

- `main.go`: Contains the Go source code for the Lambda function.
- `deployment.zip`: A ZIP file containing the compiled Go binary and its dependencies for deployment to AWS Lambda.
- `README.md`: Provides instructions on how to set up and deploy the Lambda function.

# Prerequisites

To set up and deploy the Lambda function locally, you need the following prerequisites:

- Go installed and properly configured on your local machine.
- An AWS account.
- The AWS CLI installed and configured on your local machine.

# Deployment

The deployment process involves the following steps:

1. Clone the repository.
2. Write the Lambda function code.
3. Build the Go binary.
4. Package the code into a ZIP file.
5. Create an AWS Lambda function.
6. Upload the code to the Lambda function.
7. Set up an API Gateway trigger for the Lambda function.
8. Deploy the API.
9. Test the Lambda function and API endpoint.

## Getting Started

Follow the steps below to set up and deploy the Lambda function.

### Step 1: Clone the Repository

Clone this repository to your local machine using the following command:

> ``git clone https://github.com/ahmedMHasan/weather-status-aws-lambda-function.git``


### Step 2: Write the Lambda Function

Create a new file called `main.go` and copy the code from the example into it.

### Step 3: Build the Go Binary

Open a terminal and navigate to the project directory. Run the following command to build the Go binary:

> `go build -o main.exe`


### Step 4: Package the Code

Create a ZIP file containing the binary and its dependencies. Zip Files name : `deployment.zip`


### Step 5: Create an AWS Lambda Function

1. Go to the AWS Management Console and navigate to the Lambda service.
2. Click on "Create function" and choose "Author from scratch."
3. Provide a name for your function and select "Go 1.x" as the runtime.
4. Under "Permissions," choose "Create a new role with basic Lambda permissions."
5. Click on "Create function."

### Step 6: Upload the Code

1. Scroll down to the "Function code" section.
2. Choose "Upload a .zip file" from the "Code entry type" dropdown.
3. Click on "Upload" and select the `deployment.zip` file you created earlier.
4. Set the "Handler" field to the name of your Go binary (e.g., `main`).

### Step 7: Set up an API Gateway Trigger

1. In the Lambda function configuration page, click on "Add trigger."
2. Select "API Gateway" as the trigger type.
3. Choose "Create a new API" and provide a name for your API.
4. Under "Security," select "Open" or configure authentication as desired.
5. Click on "Add" to add the trigger.

### Step 8: Deploy the API

1. In the API Gateway configuration page, click on "Actions" and select "Deploy API."
2. Choose a stage name and click on "Deploy."

### Step 9: Test the Lambda Function

1. In your Lambda function's configuration page, click on "Test."
2. Create a new test event or use the default event template.
3. Click on "Test" to execute the Lambda function.

### Step 10: Test the API Endpoint

1. In the API Gateway configuration page, go to the "Dashboard" tab.
2. Copy the "Invoke URL" for your API.
3. Use Postman or any other HTTP client to send a GET request to the API endpoint.
4. You should receive a response with the dummy weather information and current time.

Congratulations! You have successfully set up and deployed the AWS Lambda function for dummy weather status and current time using Go.

## License

This project is licensed under the [MIT License](LICENSE).
