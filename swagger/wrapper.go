 
package swagger 
import "os" 
func GetAsset() func (name string) ([]byte, error)  {return Asset} 
func GetAssetInfo() func(name string) (os.FileInfo, error) {return AssetInfo} 
func GetAssetDir() func(name string) ([]string, error) {return AssetDir}
