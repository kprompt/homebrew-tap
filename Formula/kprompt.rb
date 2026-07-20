# typed: false
# frozen_string_literal: true

class Kprompt < Formula
  desc "AI CLI to control Kubernetes with natural language"
  homepage "https://kprompt.ai"
  version "0.4.0"
  license "Apache-2.0"

  on_macos do
    on_arm do
      url "https://github.com/kprompt/kprompt/releases/download/v0.4.0/kprompt_0.4.0_darwin_arm64.tar.gz"
      sha256 "7e1302d929af5c25fd5b529608cf61865ddecfd710fba1aadbbc9e13e88df268"
    end
    on_intel do
      url "https://github.com/kprompt/kprompt/releases/download/v0.4.0/kprompt_0.4.0_darwin_amd64.tar.gz"
      sha256 "a4d0ffb34d993cee1b49d494bc74ca2f9f82e55cae69c3db930a018b91c9509b"
    end
  end

  on_linux do
    on_arm do
      url "https://github.com/kprompt/kprompt/releases/download/v0.4.0/kprompt_0.4.0_linux_arm64.tar.gz"
      sha256 "39dc3553596a562f7389dda03b5c06e3164ce210ca5829cd5fd71dbe937c36ee"
    end
    on_intel do
      url "https://github.com/kprompt/kprompt/releases/download/v0.4.0/kprompt_0.4.0_linux_amd64.tar.gz"
      sha256 "31432b9e240b673d779fcf4546cb5727efb593904beb11b028b23a5101d150f3"
    end
  end

  def install
    bin.install "kprompt"
  end

  test do
    assert_match version.to_s, shell_output("#{bin}/kprompt version")
  end
end
