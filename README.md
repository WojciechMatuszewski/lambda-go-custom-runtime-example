# Custom Go Lambda custom runtime example

This is an example of how one might write, test and deploy custom _AWS Lambda_ runtime written in Go.
Please note that the runtime itself is very simple and should not be used in any production setting.

## Local testing

I've decided to use the `local-interface-emulator` + _Docker_ setup to ensure that the _runtime loop_ I've implemented actually works.

To test things locally:

1. Build the container

   ```shell
   docker build -t hello .
   ```

2. Run the container

   ```shell
   docker run -p 9000:8080 hello
   ```

3. In another terminal, send a request to the container

   ```shell
   curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{}'
   ```

The response should be a plain string of `it works`

## Deploying

This repo uses AWS SAM for deployment

To deploy the function to AWS:

1. Build the artifacts

   ```shell
   sam build
   ```

2. Deploy the app

   ```shell
   sam deploy --guided
   ```

3. Go to AWS console and test the function

## FAQ

Q: Why did not use the `sam local invoke` for testing?
A: The `sam local invoke` invokes the function once. While it is a valid way of testing things, if I were to invoke the function once, the runtime loop would not have been tested.
