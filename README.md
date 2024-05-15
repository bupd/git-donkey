                     
<h1 align="center" style="font-weight: bold;">git-donkey 🐴</h1>


<p align="center">A Donkey Don to help you maintain your local git branch updated always.</p>


 
<h2 id="started">Introduction:</h2>

git-donkey is a tool designed to help developers keep their local git branches synchronized with the remote repository. It automates the process of fetching updates and merging them into the current branch, ensuring that your local branch is always up-to-date.
 

<h2 id="started">🚀 Getting started</h2>

<h3>Prerequisites</h3>

Please make sure you installed golang.

- [Golang](https://go.dev/doc/install)
 
<h3>Installation</h3>

To install git-donkey, you need to have Go installed on your system. Then, follow these steps:

```sh
go install github.com/bupd/git-donkey@latest
```
 
<h3>Usage</h3>

Once installed, you can use git-donkey with the following command:

```sh
git-donkey scan
```
`scan` command scans all of your local directories for github repos.
 
 <h2 id="technologies">💻 Technologies</h2>

- Golang
- cobra cli
- gogit


## Contributing
Contributions are welcome! To contribute:

1. [Fork the repository](https://github.com/bupd/git-donkey/fork).
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -am 'Add new feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Open a Pull Request.
Please ensure that your code adheres to the project's coding standards and includes appropriate tests.

## License
This project is licensed under the MIT License. See the [LICENSE](https://github.com/bupd/git-donkey/blob/main/LICENSE) file for details.
 
<h3>Documentations that might help</h3>

[📝 How to create a Pull Request](https://www.atlassian.com/br/git/tutorials/making-a-pull-request)

[💾 Commit pattern](https://gist.github.com/joshbuchea/6f47e86d2510bce28f8e7f42ae84c716)

