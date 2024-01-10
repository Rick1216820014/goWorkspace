<template>
    <div class="gvb_admin">
        <aside>
            <div>
                
            </div>
        </aside>
        <div class="main">
            <header>
                <div class="left">
                        <a-breadcrumb>
                            <a-breadcrumb-item>首页</a-breadcrumb-item>
                            <a-breadcrumb-item><a href="#">个人中心</a></a-breadcrumb-item>
                            <a-breadcrumb-item>用户列表</a-breadcrumb-item>
                            
                        </a-breadcrumb>
                </div>
                <div class="right">
                    
                        <div class="icon_actions">
                            <i class="fa fa-home"></i>
                            <i class="fa fa-moon-o"></i>
                            <i class="fa fa-sun-o"></i>
                            <i class="fa fa-arrows-alt"></i>
                        </div>
                        <div class="avatar">
                            <img src="D:\workSpace\goWorkspace\gvb\gvb_admin\src\img\微信图片_20231228094716.jpg" alt="">
                        </div>
                        <div class="drop_menu">
                            <a-dropdown placement="bottomRight">
                                <a class="ant-dropdown-link" @click.prevent>
                                高瑞克
                                    <i class="fa fa-chevron-down"></i>
                                </a>
                                <template #overlay>
                                <a-menu>
                                    <a-menu-item>
                                    <a href="javascript:;">个人中心</a>
                                    </a-menu-item>
                                    <a-menu-item>
                                    <a href="javascript:;">我的消息</a>
                                    </a-menu-item>
                                    <a-menu-item>
                                    <a href="javascript:;">文章列表</a>
                                    </a-menu-item>
                                    <a-menu-item>
                                    <a href="javascript:;">注销退出</a>
                                    </a-menu-item>
                                </a-menu>
                                </template>
                            </a-dropdown>
                        </div>
                        
                    
                </div>
                <!-- <i class="iconfont icon-yonghu"></i>
                <i class="fa fa-hand-o-up"></i> -->
            </header>
            <div class="tabs">

            </div>
            <main>

            </main>
        </div>
    </div>
    
</template>

<script setup>
    
</script>

<!-- <style lang="scss" scoped> -->
<style lang="scss" >
.gvb_admin{
    width: 100%;
    display: flex;
    aside{
        width: 240px;
        height: 100vh;
        background-color: #2b3539;
    }
    .main{
      width: calc(100% - 240px);
      
    //   position: relative; /* 添加相对定位 */
        header{
                     
            height: 60px;
            background-color: white;
            padding: 0 20px;
            display: flex;  
            justify-content: space-between;
            align-items: center;

            .right{
            display: flex;
            align-items: center; /* 垂直居中内容 */
                
            }
            
                 
            .icon_actions{
                margin-right: 20px;

                i{
                    margin-left: 10px;
                    cursor: pointer;
                    font-size: 16px;
                    color: var(--text);
                }
                
                i:hover{
                    color: var(--active);
                }
            }
            .avatar{
                
                img{
                    
                   width: 40px;
                   height: 40px;
                   border-radius: 50%; 
                }
            }
            .drop_menu{
                margin-left: 10px;
                color: var(--text);
            }
        }
        .tabs{
            height: 30px;
            border: 1px solid #f0eeee;
            
        }
        main{
            background-color: #f0eeee;
            height: calc(100vh - 90px);
        }
    }
}
</style>