# 学习感悟

K8s是一个编排容器的工具，其实也是管理应用的全生命周期的一个工具，从创建应用，应用的部署，应用提供服务，扩容缩容应用，应用更新，都非常的方便，而且可以做到故障自愈，例如一个服务器挂了，可以自动将这个服务器上的服务调度到另外一个主机上进行运行，无需进行人工干涉。

Kubernetes（K8S）是谷歌的第三个容器管理系统（前两个：Borg、Omega，这两个是谷歌内部系统，k8s是开源的），Kubernetes在Docker技术之上，为容器化的应用提供了资源调度、部署运行、服务发现和扩容缩容等丰富多样的功能。在项目公开后不久，微软、IBM、VMware、Docker、CoreOS以及SaltStack等多家公司便纷纷加入了Kubernetes社区，为该项目发展作出贡献。

k8s可以更快的更新新版本，打包应用，更新的时候可以做到不用中断服务，服务器故障不用停机，从开发环境到测试环境到生产环境的迁移极其方便，一个配置文件搞定，一次生成image，到处运行。

k8s能干什么呢？

自动化容器的部署和复制随时扩展或收缩容器规模将容器组织成组，并且提供容器间的负载均衡很容易地升级应用程序容器的新版本提供容器弹性，如果容器失效就替换它K8S有两个分类

(1)自主式pod:自我管理的，创建以后，任然是提交给API Server,由API Server接收后借助Scheduler将其调度到指定的node节点，然后由node启动此节点，如果有pod中容器出现问题了，需要重启容器，那是由Keepalived来完成，但是如果节点故障了，pod也就消失了，所以，它有这样的坏处，没办法实现全局的调度，所以建议使用第二种pod

(2)控制器管理的pod，正是控制器这种机制的使用，使得在K8S上的集群设计中，pod完全就可以叫做有生命周期的对象，而后由调度器调度到集群中的某节点，运行以后，它的任务终止也就随着被删除掉，但是有一些任务，大家知道传统上有nginx、tomcat，它们是做为守护进程运行的，那么这种运行为pod容器的话，我们要确保它是时刻运行的，一旦出现故障我们要第一时间发现，要么是重建取代它，要么是重启，那么pod控制器就可以观测并实现。
