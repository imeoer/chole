<style>
  .editor {
    overflow: auto;
    .editor-wrap {
      @mixin shadow;
      padding: 10px;
      background-color: #1e1e1e;
      border-radius: 7px;
      margin-bottom: 20px;
    }
    #editor {
      width: 486px;
      height: 250px;
      .monaco-editor .view-line {
        border-bottom: 1px dashed #313131;
      }
    }
    .buttons {
      display: flex;
      float: right;
      .apply {
        margin-left: 30px;
      }
    }
  }
</style>

<template>
  <div class="editor box" transition="show">
    <div class="editor-wrap">
      <div id="editor"></div>
    </div>
    <div class="buttons">
      <circle-button class="cancel grey" icon="close" v-link="'/'">取消</circle-button>
      <circle-button class="apply green" icon="check">应用</circle-button>
    </div>
  </div>
</template>

<script>
import CircleButton from './common/CircleButton.vue'

export default {
  components: {
    CircleButton
  },
  vuex: {
    getters: {
    }
  },
  ready() {
    window.load(['vs/editor/editor.main'], () => {
      this.editor = window.monaco.editor.create(document.getElementById('editor'), {
        value: [
          'rules:',
          '  web1:',
          '    out: 8001',
          '    in: 8000'
        ].join('\n'),
        fontFamily: 'monaco',
        fontWeight: 'lighter',
        fontSize: 12,
        theme: 'vs-dark',
        // language: 'javascript',
        cursorBlinking: 'smooth',
        scrollbar: {
          horizontalScrollbarSize: 5,
          verticalScrollbarSize: 5
        },
        wordBasedSuggestions: true,
        quickSuggestions: true,
        lineNumbers: false,
        revealHorizontalRightPadding: 0,
        lineDecorationsWidth: 0,
        overviewRulerLanes: 0,
        roundedSelection: false
      })
    })
  },
  beforeDestroy() {
    this.editor.destroy()
  }
}
</script>
