#!/usr/bin/env bash

package=$1
output_dir=$2
targets=$3

if [[ -z "$package" || -z "$output_dir" || -z "$targets" ]]; then
  echo "usage: $0 <package-name> <output_dir> <platform/arch/version|another_platform/arch/version>"
  echo "example: ./go_multi_arch_build.sh github.com/someone/nice-package /absolute/binary/output/path 'windows/amd64|linux/amd64|linux/arm/5'"
  exit 1
fi

package_split=(${package//\// })
package_name=${package_split[-1]}
platforms=(${targets//|// })

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name=$package_name'-'$GOOS'-'$GOARCH
    if [ $GOARCH = "arm" ]; then
        GOARM=${platform_split[2]}
        output_name+="-$GOARM"
    fi
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi

    if [ $GOARCH = "arm" ]; then
        env GOOS=$GOOS GOARCH=$GOARCH GOARM=$GOARM go build -o $output_dir/$output_name $package
    else
        env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_dir/$output_name $package
    fi

    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done