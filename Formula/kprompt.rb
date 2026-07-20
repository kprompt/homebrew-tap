# typed: false
# frozen_string_literal: true

class Kprompt < Formula
  desc "AI CLI to control Kubernetes with natural language"
  homepage "https://kprompt.ai"
  version "0.3.0"
  license "Apache-2.0"

  on_macos do
    on_arm do
      url "https://github.com/kprompt/kprompt/releases/download/v0.3.0/kprompt_0.3.0_darwin_arm64.tar.gz"
      sha256 "df50c6af86e0fc84f00c278df55394a8cfdf3158f745249c2ec6d9fe3abbc1fc"
    end
    on_intel do
      url "https://github.com/kprompt/kprompt/releases/download/v0.3.0/kprompt_0.3.0_darwin_amd64.tar.gz"
      sha256 "7acf30abb432f40cfb756b2b4f403d0ccc167abc7b44fd1aa07a803380ac8ff1"
    end
  end

  on_linux do
    on_arm do
      url "https://github.com/kprompt/kprompt/releases/download/v0.3.0/kprompt_0.3.0_linux_arm64.tar.gz"
      sha256 "54ad78ba22ac95c2b2d7e89797517ea72a3301a7fcc6cb25ec116e2fa53f037c"
    end
    on_intel do
      url "https://github.com/kprompt/kprompt/releases/download/v0.3.0/kprompt_0.3.0_linux_amd64.tar.gz"
      sha256 "f530b44cb1e61db0087d853ec3234870f99bba0808ea35e4dfbf0144d3052023"
    end
  end

  def install
    bin.install "kprompt"
  end

  test do
    assert_match version.to_s, shell_output("#{bin}/kprompt version")
  end
end
