# typed: false
# frozen_string_literal: true

class Kprompt < Formula
  desc "AI Kubernetes CLI: natural language to a reviewable plan, then approve before apply"
  homepage "https://kprompt.ai"
  version "0.11.0"
  license "Apache-2.0"

  on_macos do
    on_arm do
      url "https://github.com/kprompt/kprompt/releases/download/v0.11.0/kprompt_0.11.0_darwin_arm64.tar.gz"
      sha256 "ffff6b5c1d77ba05d83fd864265a0746005a59b97335503c711bd8d985d174d8"
    end
    on_intel do
      url "https://github.com/kprompt/kprompt/releases/download/v0.11.0/kprompt_0.11.0_darwin_amd64.tar.gz"
      sha256 "2e02d81f1e4983f21cbab021a0526655d8a1ac589bc4921ad1b23060edef3b9c"
    end
  end

  on_linux do
    on_arm do
      url "https://github.com/kprompt/kprompt/releases/download/v0.11.0/kprompt_0.11.0_linux_arm64.tar.gz"
      sha256 "3b775cf290c199b302fd223d52dc22be26386ef10da59f7bb89a908ead5f037d"
    end
    on_intel do
      url "https://github.com/kprompt/kprompt/releases/download/v0.11.0/kprompt_0.11.0_linux_amd64.tar.gz"
      sha256 "2a4fa7ed182e22769f501764e4642db2eb609224ef57deeb0b209ef9b12969b8"
    end
  end

  def install
    bin.install "kprompt"
  end

  test do
    assert_match version.to_s, shell_output("#{bin}/kprompt version")
  end
end
