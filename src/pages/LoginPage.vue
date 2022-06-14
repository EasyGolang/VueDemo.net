<script setup lang="ts"></script>

<template>
  <main>
    <h1>登录</h1>
    <input type="text" name="Email" placeholder="请输入邮箱" autoComplete="new-password" />
    <br />
    <input type="password" name="Password" placeholder="请输入密码" autoComplete="new-password" />
    <br />
    <br />
    Email: {{ Email }}
    <br />
    <br />
    Password: {{ Password }}
    <br />
    <br />

    <button on:click="{Submit}">提交</button>
    <br />

    返回信息：
    <div v-if="resMsg" class="resMsg">{{ JSON.stringify(resMsg) }}</div>
  </main>
</template>

<style lang="less" scoped>
.resMsg {
  margin: 12px;
  border: 10px solid gray;
  padding: 15px;
  &.success {
    border-color: green;
  }
  &.fail {
    border-color: red;
  }
}
</style>
