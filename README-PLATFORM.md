# Go Platform Challenge

## Instructions

We think infrastructure is best represented as code, and provisioning of resources should be automated as much as possible. Build out some basic Infrastructure in GCP to deploy our [Application](application/) that can be used in a repeatable way. 

**Branch out** this repository, and implement the scenario described below.

* You *shall* complete this challenge using **Terraform**,
* You *shall* complete this challenege using **Kubernetes**,
* You *should* try to show your platform process to present a **production-ready** infrastructure,
* You *should* work in infrastructure folder, so all your artefacts ends up in `infrastructure` folder,
* You *should* provide a diagrma explaining your infrastructure architecture
* Please, describe your approach, your technical choices, including architectural, and anything you want us to know in a `Results.md` file,

## The Scenario

Your mission is to deploy **microservice** implemented by your Software Engineer.

* the endpoint is exposed at `/api/v1/applications`
* the endpoint allows to manage data from a database (see explanations below)
* the query may takes parameters, to filter the results. You are expected to propose any kind of query params you may find useful to provide a rich user experience,
* the endpoint returns a JSON response with an array of matching results

## Application 

Use a tool of your choice such as Bash, Github Actions or similar to automate the docker build and deploy of the [Application](application/) that serves some static or dynamic content. 

* Using Containers as part of your automation.
* Creating a CI pipeline, using a tool of your choice, that build the service.
* Creating a CD pipilien, using a tool of your choice, deploying to GCP cloud environment
* Serve traffic from 443 port with self-signed certificate would be highly appreciate.

## The database

You are expected to integrate the data a **mongodb** (container)

## How we review ?

Your application will be reviewed by one of our team member. We do take into consideration your experience level.

**We value quality over feature-completeness**. The goal of this code sample is to help us identify what you consider **production-ready** code. You should consider this code ready for final review with your colleague, i.e. this would be the last step before deploying to production.

The aspects of your code we will assess include:

* **Architecture**: how well it is a architected? does the code make use of design patterns ? are the DevOps/DevSecOps principles followed ?
* **Clarity**: does the documentation (result.md) clearly and concisely explains the problem and solution? Are technical tradeoffs explained ?
* **Correctness**: does the code do what was asked ? If there is anything missing, does the result.md explain why it is missing?
* **Code quality**: is the code simple, easy to understand, and maintainable ? Are there any code smells or other red flags ? Does it follows coding principles such as the single responsibility principle ? Is the coding style consistent with the language's guidelines ? Is it consistent throughout the codebase ?
* **Security**: are there any obvious vulnerability ?
* **Testing**: Will they be difficult to change if the requirements of the application were to change ? 
* **Technical choices**: do choices of tools, patterns, architecture etc. seem appropriate for the chosen application ?
* **Documentation**: is your platform documented ? With a decent explanation ?
* **Production-readiness**: does the infra include monitoring ? logging ? proper error handling ?

Bonus point (those items are optional):

* **Scalability**: will technical choices scale well ? If not, is there a discussion of those choices in the result.md ?

## Notes

* We're evaluating you on completion of the scenario, but also on style. Do you commit often? Did you make code that you would enjoy reading ? Things like that are important to our team 👊
* We expect you to provide a code as you would do in any professional situation, with future maintenability and continuous deployment constraints in mind, with deployment or installation instructions, ...
* This should be fun! If you're stuck on an instruction or if something needs clarification, you can start discussion on the repository board.

## Team Work
Work along your software engineer and make sure you are aligned with them, and all your decisions are communicated with them

## Done ?

Great job! When you're all finished, **make sure it is already communicated with your software engineer and it is implemented as per specification**, open a PR and we'll review it together 🙌


## Thanks

* http://onqlave.com
