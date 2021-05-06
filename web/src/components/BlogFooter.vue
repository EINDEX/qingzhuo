<template>
  <footer>
    <div
      class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 pt-6 pb-4 items-centers justify-items-center"
    >
      <div class="grid h-24 md:grid-cols-4 sm:grid-cols-1 gap-4 justify-items-center my-4">
        <div class="md:col-start-4 sm:col-auto">
          <div class="flex flex-row gap-4">
            <img class="h-24 rounded-2xl" src="/@/assets/avatar.png" />
            <div class="col-start-4">
              <div class="flex flex-col items-start justify-around leading-6 h-full">
                <a
                  href="https://github.com/eindex"
                  rel="noreferrer"
                  target="_blank"
                  class="flex align-middle flex-row h-6 gap-1"
                >
                  <img class="h-5 w-5" src="/@/assets/github.svg" />
                  Github
                </a>
                <a
                  href="mailto:eindex.me@outlook.com"
                  target="_blank"
                  class="flex align-middle flex-row h-6 gap-1"
                >
                  <img class="h-6 w-5" src="/@/assets/icon-mail.svg" />
                  Email
                </a>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="grid justify-items-center">
        <a
          class="flex flex-row h-4 leading-4 gap-1 text-gray-500"
          href="http://www.beian.gov.cn/portal/index"
          rel="noreferrer"
          target="_blank"
          ><img class="h-4 w-4" src="/@/assets/gongan.png" />鄂ICP备15018232号-1</a
        >
      </div>
    </div>
  </footer>
</template>
