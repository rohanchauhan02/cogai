# CogAI (Cognitive AI)

**CogAI** is an AI-powered CLI tool designed to streamline your Go development experience. Built specifically for GoLang enthusiasts, CogniTerm assists developers by automating repository creation, package management, and project structuring. Leveraging the power of AI, it intelligently suggests and implements Go packages based on your project needs, ensuring optimized performance and clean architecture.

With **CogAI**, you can:

- Effortlessly create repositories and manage dependencies.
- Get smart suggestions for packages tailored to your project requirements.
- Automate repetitive tasks, allowing you to focus on writing high-quality Go code.
- Benefit from AI-driven insights for enhancing code quality and efficiency.

Whether you're a seasoned Go developer or just starting, **CogAI** provides the support you need to accelerate development with minimal effort.

## How to Use

### Installation

To install CogAI, simply run the following command:

```bash
go install github.com/cogai/cogai@latest
```

### Usage

To use CogAI, simply run the following command:

```bash
cogai <command>
```

### Commands

#### `Initiate CogAI`

- `cogai init`: Initiate cogai.

![CogAI](https://github.com/rohanchauhan02/cogai/blob/master/docs/logo.png)

#### `Get disk info`

- `info du`: Get disk info.

#### `Env Management`

- `set OPEN_AI_KEY=XXXXXXX`: Add openai key to env file.
- `delete -k/--key open_ai_key(or OPEN_AI_KEY)`: Delete the key from env file.

- `get -a/--all`: List all keys in env file.
- `get -k/--key open_ai_key(or OPEN_AI_KEY)`: Get the value of the key.
