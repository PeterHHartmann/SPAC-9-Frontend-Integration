import { AppSidebar } from '@/components/layout/AppSidebar';
import { PageLayout } from '@/components/layout/PageLayout';
import { SidebarInset, SidebarProvider } from '@/components/ui/sidebar';
import { Outlet } from 'react-router-dom';

export const RootLayout: React.FC = () => (
	<SidebarProvider>
		<AppSidebar />
		<SidebarInset>
			<PageLayout>
				{/* Outlet component is where route components will be mounted */}
				<Outlet />
			</PageLayout>
		</SidebarInset>
	</SidebarProvider >
);