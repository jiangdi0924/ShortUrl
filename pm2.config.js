module.exports = {
  apps: [
    {
      name: "gofiber",
      script: "./main",   // script当成配置文件，在go中通过os.Args[1]获取到
      instances: 1,
      exec_mode: "fork",    // 一定要是fork
      interpreter: "none",   // windows下加.exe
      env: {              // 环境变量
        myenv: "product",
      },
    }
  ]
}