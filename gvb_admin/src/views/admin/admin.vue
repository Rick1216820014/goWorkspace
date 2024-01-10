<template>
    <div class="gvb_admin">
        <aside>
            <div>
                <i class="iconfont icon-yonghu"></i>
            </div>
        </aside>
        <div class="main">
            <header>

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
      position: relative; /* 添加相对定位 */
        header{
            height: 60px;
            background-color: white;
            // position: absolute; /* 添加绝对定位 */
            // top: 0; /* 将头部置于最上方 */
            // left: 0; /* 左侧与父级容器对齐 */
            // right: 0; /* 右侧与父级容器对齐 */
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