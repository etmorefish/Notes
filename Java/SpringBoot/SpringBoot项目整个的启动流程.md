# SpringBoot项目整个的启动流程

> @SpringBootApplication 是一个复合注解，它由三个注解组成。

- @SpringBootConfiguration（打开是@Configuration）：表明这是一个 Java 配置类。
- @EnableAutoConfiguration：自动配置注解，会将所有符合自动配置条件的@Configuration 配置加载到 IOC 容器。
- @ComponentScan：扫描注解，自动扫描符合条件的组件（@Service、@Component）或者 bean 定义，记载到 IOC 容器中。

## spring boot 启动流程

1. 从 spring.factories 配置文件中加载 EventPublishingRunListener 对象，该对象拥有 SimpleApplicationEventMulticaster 属性，即在 SpringBoot 启动过程的不同阶段用来发射内置的生命周期事件;

2. 准备环境变量，包括系统变量，环境变量，命令行参数，默认变量，servlet 相关配置变量，随机值以及配置文件（比如 application.properties）等;

3. 控制台打印 SpringBoot 的 bannner 标志；

4. 根据不同类型环境创建不同类型的 applicationcontext 容器，如果是 servlet 环境，创建的就是 AnnotationConfigServletWebServerApplicationContext 容器对象；

5. 从 spring.factories 配置文件中加载 FailureAnalyzers 对象,用来报告 SpringBoot 启动过程中的异常；

6. 为刚创建的容器对象做一些初始化工作，准备一些容器属性值等，对 ApplicationContext 应用一些相关的后置处理和调用各个 ApplicationContextInitializer 的初始化方法来执行一些初始化逻辑等；

7. 刷新容器，这一步至关重要。比如调用 bean factory 的后置处理器，注册 BeanPostProcessor 后置处理器，初始化事件广播器且广播事件，初始化剩下的单例 bean 和 SpringBoot 创建内嵌的 Tomcat 服务器等等重要且复杂的逻辑都在这里实现，主要步骤可见代码的注释，关于这里的逻辑会在以后的 spring 源码分析专题详细分析；

8. 执行刷新容器后的后置处理逻辑，注意这里为空方法；

9. 调用 ApplicationRunner 和 CommandLineRunner 的 run 方法，我们实现这两个接口可以在 spring 容器启动后需要的一些东西比如加载一些业务数据等;

10. 报告启动异常，即若启动过程中抛出异常，此时用 FailureAnalyzers 来报告异常;

11. 最终返回容器对象，这里调用方法没有声明对象来接收。

